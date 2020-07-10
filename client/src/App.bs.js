'use strict';

var React = require("react");
var Caml_option = require("bs-platform/lib/js/caml_option.js");
var Navbar$Client = require("./Navbar.bs.js");
var MaterialUi_Grid = require("@jsiebern/bs-material-ui/src/MaterialUi_Grid.bs.js");
var Core = require("@material-ui/core");
var MaterialUi_CssBaseline = require("@jsiebern/bs-material-ui/src/MaterialUi_CssBaseline.bs.js");

function App(Props) {
  return React.createElement(React.Fragment, undefined, React.createElement(Core.CssBaseline, MaterialUi_CssBaseline.makeProps(undefined, undefined, undefined, undefined, undefined, undefined, undefined)), React.createElement(Navbar$Client.make, {}), React.createElement(Core.Grid, MaterialUi_Grid.makeProps(undefined, undefined, Caml_option.some(null), undefined, undefined, true, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined), React.createElement(Core.Grid, MaterialUi_Grid.makeProps(undefined, undefined, "Grid item", undefined, undefined, undefined, undefined, true, undefined, undefined, undefined, undefined, undefined, undefined, undefined, /* V2 */19228, undefined, undefined, undefined, undefined, undefined, undefined, undefined)), React.createElement(Core.Grid, MaterialUi_Grid.makeProps(undefined, undefined, "Another Grid item", undefined, undefined, undefined, undefined, true, undefined, undefined, undefined, undefined, undefined, undefined, undefined, /* V3 */19229, undefined, undefined, undefined, undefined, undefined, undefined, undefined))));
}

var make = App;

exports.make = make;
/* react Not a pure module */
