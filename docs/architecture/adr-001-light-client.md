# ADR 001: Light Client

## Changelog

* 23.06.24: Initial draft
* 12.07.24: Update with Custom light client using IBC packet commitment proofs instead of full state proofs solution

## Status

DRAFT

## Abstract

This ADR proposes a solution for verifying the behavior of optimistic rollups within the Cosmos ecosystem until 
state proofs become widely available. 
Optimistic rollups face challenges with interoperability due to a dispute window that delays state finality. 

The Pessimistic Rollup project addresses this by enabling Proof of Stake Cosmos SDK-based blockchains to verify rollup 
states using validators who run full nodes of the rollup. The initial prototype used two light clients, 
but it required adjustments for clarity and complexity reduction. 

The chosen solution is a conditional light client with a prover light client, which maintains modularity, 
leverages existing light client functionality, and aligns closely with IBC design. 
This approach avoids confusion, simplifies relayer adjustments, 
and allows future extensibility for different proving mechanisms.

## Context

The modular thesis is expected to bring many new rollups to the Cosmos ecosystem.
Until state proofs (ZK) are a viable option, we will see many (most?) of the rollups
using so-called optimistic rollups. State proofs are undoubtedly the future for rollups,
but even with all the advancements in ZK tech, provable state transitions will take time
to be available everywhere.

Optimistic rollups have an inherent problem with interoperability, since the state of a rollup
is not considered final/safe until a dispute window (a period of time where someone can 
provide a fraud proof) - of typically 7 days - has passed.

The Pessimistic Rollup project was conceived to solve this problem by allowing any Proof of Stake
Cosmos SDK-based blockchain to leverage its economic security and verify that a given rollup
is behaving correctly. The validators do this by running a full node of the rollup and report
the state and height back to its own chain. With this information a chain can safely accept and verify
IBC packets from the rollup.

The initial prototype, developed during the Celestia Infinite Space Bazaar hackathon, used
two light clients (standard tendermint + a custom light client updated with ABCI++ Vote Extensions)
and a custom Cosmos SDK module to achieve Pessimistic Validation for a rollup. The design was
functional, but had some downsides that makes adjustments necessary. In particular, the design
required a standard tendermint light client to be running and kept alive, which could be used directly by anyone.
While the denoms coming across a connection with that light client would not be canonical, it was
an unnecessary confusing to have an operational light client that should not be used. It could
have been solved by requiring an IBC middleware to stop packets or channel creation, but this
would require more complexity for integrators and the protocol itself.

## Alternatives

### Custom conditional light client

This is the solution outlined in the initial prototype. The custom light client would contain the
pessimistic validation, and depend on a standard light client to provide inclusion proofs.

![Initial solution diagram](initial-solution.png)

The main benefits of this solution are:
- A lot of the logic is already written for this
- It requires no modification to existing light clients (i.e. you can use 07-tendermint as is)

The main drawbacks of this solution are:
- The standard light client is still running and can be used by anyone, which is confusing
- Relayers would have to be adjusted to support the custom light client

### IBC Middleware

A simpler solution to the problem (in terms of moving pieces at least) would be to not have any 
new light clients, and instead use IBC Middleware to filter packets.

![IBC Middleware solution diagram](ibc-middleware-solution.png)

IBC Middleware lives between the core IBC modules (client, channel, connection) and the application
modules (ICS20, ICS721, etc). It allows you to manipulate or even block packets as they are coming in.

If Pessimistic Validation was implemented as an IBC middleware, you could use existing light clients
(such as 07-tendermint) without any modifications. The middleware would keep track of the validation
updates of the counterparty rollup, and only let through packets that are less than or equal to the
latest height proven to be safe.

Unfortunately, IBC middleware does not have direct access to the underlying light client or the provided commitment proof, so it would
have to be able to have a mechanism to find or construct all the above. It might be possible, but would be quite hacky and
probably go against both the design and flow of IBC. 

The main benefits of this solution are:
- No new light clients need to be implemented
- Less to implement in general
- Existing light clients can be used without any modification
- Potential to implement other filtering mechanisms, such as an escrow system like eIBC

An IBC middleware solution could implement a fee-taking system, where a small fee is redirected from ICS20 packets.
This, however, would only work for ICS20 packets, and would not help with other types of IBC packets.

While this solution is materially simpler than any that involves custom light clients, it also has some
drawbacks that are important to note:
- Relayers and clients would not be able to check that a light client has been proven by the pessimistic validators, 
  because all of that would happen in the middleware. If you simply queried the light client you would think that a packet can be proven.
- If even possible, it would be hacky and a counterintuitive flow

### Conditional Light Client with a prover light client

The solution with a conditional light client is based on a concept discussed in
[https://github.com/cosmos/ibc-go/issues/5112](ibc-go issue 5112). The idea is to have a light client
that queries another light client for some verification it cannot do itself. In the issue the main topic
was data inclusion from the DA layer, but in this case we are applying it to the Pessimistic Validation.

The idea is to have an almost standard light client that (such as the tendermint light client) that simply
queries one or more prover light clients. The prover light client would be a custom light client that can
implement different types of proving mechanisms. In the case of Pessimistic Validation, the prover light client
would provide proof in the form of a trusted committee (validators with stake) signature that the rollup is behaving correctly.

In the future, this could be extended to provide other proving mechanisms, such as ZK proofs, DA inclusion proofs, etc.

The prover light clients are updated by the validators of the chain it is running on, and would not use relayers
for client state updates. This would happen through another mechanism which will be described in another ADR.

![Conditional Light Client with prover solution diagram](conditional-light-client-with-prover-solution.png)

Using a conditional light client we can avoid most of the issues from the initial prototype, and get some added benefits:
- The standard light client can be used as normal by users and relayers, and there is no confusion about what the current validated height and state is (as you can query it directly)
- Relayers would not need to make any adjusments as long as they support the normal light client (e.g. 07-tendermint)
- The protocol is more modular and can be extended with different proving mechanisms
- Follows the design of IBC more closely
- Existing fee mechanisms can be used without any changes

The main drawbacks of this solution are:
- Up to two custom light clients (even though the standard light client is a fork with minimal changes)
- No ability to filter packets directly or apply other middleware-like mechanisms (e.g. eIBC)
- Every new type of chain requires a new light client and potentially a new prover light client as well

### Custom light client using IBC packet commitment proofs instead of full state proofs

This solution would be to have a custom light client that uses IBC packet commitment proofs to verify the state of the rollup.
It would be able to support any chain that can implement IBC, including EVM chains.

For updating the light client, the validators would provide signed IBC Packet Commitments that would be saved in ConsensusState.

![Updating Light Client with IBC Packet Commitments diagram](update-light-client-packet-commitments.png)

Next, when receiving a packet, the light client would verify that the packet commitment is included in the ConsensusState.

![Receiving packets in Light Client with IBC Packet Commitments diagram](receive-msg-light-client-packet-commitments.png)

The security assumptions of this solution are the same as for the other solutions because the IBC Packet Commitments are
still signed by the validators.

The main benefits of this solution are:
- It is more general and can be used for any chain that can implement IBC
- It only requires a single custom light client for all chains
- It is simpler to both integrate and maintain

The main drawbacks of this solution are:
- It is a custom light client that most likely requires some relayer adjustments
- No ability to filter packets directly or apply other middleware-like mechanisms (e.g. eIBC)

## Decision

Decision pending...

## Consequences

For Pessimistic Validation
* 

For Chains
* 

For Relayers
* 

For Users
* 

### Backwards Compatibility

Not applicable, since Pessimistic Validation is not in production yet.

## References

* [https://github.com/cosmos/ibc-go/issues/5112](Conditional clients ibc-go issue)
