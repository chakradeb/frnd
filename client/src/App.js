import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';

import routes from "./routes";
import Header from "./components/Header";

const App = function () {
  return (
      <React.Fragment>
          <Header/>
          <Router>
              {routes}
          </Router>
      </React.Fragment>
  )
};

export default App;
