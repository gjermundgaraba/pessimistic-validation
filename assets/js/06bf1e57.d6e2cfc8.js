"use strict";(self.webpackChunkinterchain_attestation_docs=self.webpackChunkinterchain_attestation_docs||[]).push([[263],{1704:(t,e,n)=>{n.r(e),n.d(e,{assets:()=>c,contentTitle:()=>r,default:()=>h,frontMatter:()=>i,metadata:()=>s,toc:()=>d});var o=n(4848),a=n(8453);const i={sidebar_position:4},r="Vote Extensions",s={id:"architecture/vote-extensions",title:"Vote Extensions",description:"ABCI++ is a CometBFT interface that allows a chain to add more functionality to the low-level parts of their application.",source:"@site/docs/architecture/vote-extensions.md",sourceDirName:"architecture",slug:"/architecture/vote-extensions",permalink:"/docs/architecture/vote-extensions",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/architecture/vote-extensions.md",tags:[],version:"current",sidebarPosition:4,frontMatter:{sidebar_position:4},sidebar:"docsSidebar",previous:{title:"Attestation Sidecar",permalink:"/docs/architecture/sidecar"},next:{title:"Attestation Light Client",permalink:"/docs/architecture/light-client"}},c={},d=[];function l(t){const e={a:"a",h1:"h1",header:"header",p:"p",...(0,a.R)(),...t.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(e.header,{children:(0,o.jsx)(e.h1,{id:"vote-extensions",children:"Vote Extensions"})}),"\n",(0,o.jsxs)(e.p,{children:["ABCI++ is a CometBFT interface that allows a chain to add more functionality to the low-level parts of their application.\nIn particular, for our purposes, it allows validators to communicate with each other to aggregate potentially disparate information\nand come to consensus on an aggregate value. To read more about ABCI++, please refer to the official documentation ",(0,o.jsx)(e.a,{href:"https://docs.cometbft.com/v0.38/spec/abci/",children:"here"}),"."]}),"\n",(0,o.jsx)(e.p,{children:"In the context of Interchain Attestation, we use the ABCI++ interface to fetch attestations from the sidecar,\naggregate them, and send them to the light client for verification and client updates."}),"\n",(0,o.jsx)(e.p,{children:"TODO: Add illustration with all the different callbacks"})]})}function h(t={}){const{wrapper:e}={...(0,a.R)(),...t.components};return e?(0,o.jsx)(e,{...t,children:(0,o.jsx)(l,{...t})}):l(t)}},8453:(t,e,n)=>{n.d(e,{R:()=>r,x:()=>s});var o=n(6540);const a={},i=o.createContext(a);function r(t){const e=o.useContext(i);return o.useMemo((function(){return"function"==typeof t?t(e):{...e,...t}}),[e,t])}function s(t){let e;return e=t.disableParentContext?"function"==typeof t.components?t.components(a):t.components||a:r(t.components),o.createElement(i.Provider,{value:e},t.children)}}}]);