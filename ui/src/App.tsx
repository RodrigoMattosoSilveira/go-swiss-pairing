import React from 'react';
import { BrowserRouter as Router, Routes, Route, NavLink} from 'react-router-dom';

import HomePage from './home/HomePage';
import MembersPage from "./members/MembersPage";
import MemberPage from "./members/MemberPage";

import './App.css';

function App() {
    // return (
    //   <div className="container">
    //          <MembersPage />
    //       </div>
    //   );}
    return (
        <Router>
            <header className="sticky">
                <span className="logo">
                    <img src="/assets/logo-3.svg" alt="logo" width="49" height="99" />
                </span>
                <NavLink to="/"  className="button rounded">
                    <span className="icon-home"></span>
                    Home
                </NavLink>
                <NavLink to="/members" className="button rounded">
                    Members
                </NavLink>
            </header>
            <div className="container">
                <Routes>
                    <Route path="/" element={<HomePage/>}/>
                    <Route path="/members" element={<MembersPage/>}/>
                    <Route path="/members/:id" element={<MemberPage />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
