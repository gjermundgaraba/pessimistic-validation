"use strict";(self.webpackChunkinterchain_attestation_docs=self.webpackChunkinterchain_attestation_docs||[]).push([[490],{2533:(e,n,i)=>{i.r(n),i.d(n,{assets:()=>r,contentTitle:()=>o,default:()=>u,frontMatter:()=>t,metadata:()=>d,toc:()=>a});var l=i(4848),s=i(8453);const t={sidebar_position:2},o="Building",d={id:"developing/building",title:"Building",description:"Building all the modules",source:"@site/docs/developing/building.md",sourceDirName:"developing",slug:"/developing/building",permalink:"/docs/developing/building",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/developing/building.md",tags:[],version:"current",sidebarPosition:2,frontMatter:{sidebar_position:2},sidebar:"docsSidebar",previous:{title:"Prerequisites",permalink:"/docs/developing/prerequisites"},next:{title:"Testing",permalink:"/docs/developing/testing"}},r={},a=[{value:"Building all the modules",id:"building-all-the-modules",level:2},{value:"Linting",id:"linting",level:2},{value:"Proto",id:"proto",level:2},{value:"Docker images (used for testing)",id:"docker-images-used-for-testing",level:2},{value:"Simapps",id:"simapps",level:2},{value:"Running locally",id:"running-locally",level:2}];function c(e){const n={code:"code",h1:"h1",h2:"h2",header:"header",li:"li",p:"p",pre:"pre",ul:"ul",...(0,s.R)(),...e.components};return(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)(n.header,{children:(0,l.jsx)(n.h1,{id:"building",children:"Building"})}),"\n",(0,l.jsx)(n.h2,{id:"building-all-the-modules",children:"Building all the modules"}),"\n",(0,l.jsx)(n.p,{children:"To build all the modules, you can use the following command:"}),"\n",(0,l.jsx)(n.pre,{children:(0,l.jsx)(n.code,{className:"language-bash",children:"$ just build\n"})}),"\n",(0,l.jsx)(n.h2,{id:"linting",children:"Linting"}),"\n",(0,l.jsx)(n.p,{children:"To lint all the modules, you can use the following command:"}),"\n",(0,l.jsx)(n.pre,{children:(0,l.jsx)(n.code,{className:"language-bash",children:"$ just lint\n"})}),"\n",(0,l.jsx)(n.h2,{id:"proto",children:"Proto"}),"\n",(0,l.jsx)(n.p,{children:"To generate go code for all the proto files, you can use the following command:"}),"\n",(0,l.jsx)(n.pre,{children:(0,l.jsx)(n.code,{className:"language-bash",children:"$ just proto-gen\n"})}),"\n",(0,l.jsx)(n.h2,{id:"docker-images-used-for-testing",children:"Docker images (used for testing)"}),"\n",(0,l.jsx)(n.p,{children:"There is a set of docker images used for the e2e tests. To build these images, you can use the following command:"}),"\n",(0,l.jsx)(n.pre,{children:(0,l.jsx)(n.code,{className:"language-bash",children:"$ just build-docker-images\n"})}),"\n",(0,l.jsx)(n.h2,{id:"simapps",children:"Simapps"}),"\n",(0,l.jsx)(n.p,{children:"If you want to build and install the simapp binaries locally, you can use the following command:"}),"\n",(0,l.jsx)(n.pre,{children:(0,l.jsx)(n.code,{className:"language-bash",children:"just install-simapps\n"})}),"\n",(0,l.jsx)(n.h2,{id:"running-locally",children:"Running locally"}),"\n",(0,l.jsx)(n.p,{children:"There is a command that spins up a local test environment with the following components:"}),"\n",(0,l.jsxs)(n.ul,{children:["\n",(0,l.jsx)(n.li,{children:"Simapp (Cosmos SDK chain with Interchain Attestion integrated)"}),"\n",(0,l.jsx)(n.li,{children:"Rollupsimapp (Rollkit rollup)"}),"\n",(0,l.jsx)(n.li,{children:"Mock DA service (for the rollup)"}),"\n",(0,l.jsx)(n.li,{children:"Sidecar"}),"\n",(0,l.jsxs)(n.li,{children:["Configuration","\n",(0,l.jsxs)(n.ul,{children:["\n",(0,l.jsx)(n.li,{children:"Light clients"}),"\n",(0,l.jsx)(n.li,{children:"Connections"}),"\n",(0,l.jsx)(n.li,{children:"Channels"}),"\n",(0,l.jsx)(n.li,{children:"Validator registered and wired up to Interchain Attestation"}),"\n"]}),"\n"]}),"\n"]}),"\n",(0,l.jsx)(n.p,{children:"To run this environment, you can use the following command:"}),"\n",(0,l.jsx)(n.pre,{children:(0,l.jsx)(n.code,{className:"language-bash",children:"just serve\n"})})]})}function u(e={}){const{wrapper:n}={...(0,s.R)(),...e.components};return n?(0,l.jsx)(n,{...e,children:(0,l.jsx)(c,{...e})}):c(e)}},8453:(e,n,i)=>{i.d(n,{R:()=>o,x:()=>d});var l=i(6540);const s={},t=l.createContext(s);function o(e){const n=l.useContext(t);return l.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function d(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:o(e.components),l.createElement(t.Provider,{value:n},e.children)}}}]);