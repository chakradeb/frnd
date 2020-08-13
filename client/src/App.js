import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';

import routes from "./routes"

const App = function () {
  return (
      <Router>
        {routes}
      </Router>
  )
};

export default App;
