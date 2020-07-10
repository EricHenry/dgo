'use strict';

var React = require("react");
var App$Client = require("./App.bs.js");
var ReactDOMRe = require("reason-react/src/legacy/ReactDOMRe.bs.js");

ReactDOMRe.renderToElementWithId(React.createElement(App$Client.make, {}), "root");

/*  Not a pure module */
