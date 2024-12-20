import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import LinksList from './LinksList';
import AddLink from './AddLink';

const App = () => {
    return (
        <Router>
            <nav>
                <ul>
                    <li><a href="/">Links</a></li>
                    <li><a href="/add">Add Link</a></li>
                </ul>
            </nav>

            <Switch>
                <Route path="/" exact component={LinksList} />
                <Route path="/add" component={AddLink} />
            </Switch>
        </Router>
    );
};

export default App;
