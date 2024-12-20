import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LinksList from './LinksList';
import AddLink from './AddLink';
import UpdateLink from './UpdateLink';
import './App.css';
import './button.css';
import './div.css';

const App = () => {
    return (
        <Router>
            <nav>
                <ul>
                    <li><a href="/">Links</a></li>
                    <li><a href="/add">Add Link</a></li>
                </ul>
            </nav>

            <Routes>
                <Route path="/" element={<LinksList />} />
                <Route path="/add" element={<AddLink />} />
                <Route path="/update/:id" element={<UpdateLink />} />
            </Routes>

            <ul id="author">
                <li><a href="https://www.youtube.com/watch?v=dQw4w9WgXcQ" target="_blank" rel="noreferrer"><text>DO NOT</text> Contact us       </a></li>
            </ul>
        </Router>
    );
};

export default App;
