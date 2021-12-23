import React from 'react';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import {Link, Route, Routes, useNavigate} from "react-router-dom";

import Login from "./components/login.component";
import SignUp from "./components/signup.component";
import Home from "./components/home.component";
import Dashboard from "./components/dashboard.component";
import useToken from "./components/useToken";

function clearToken() {
    sessionStorage.removeItem("token");
}

function App() {
    const { token, setToken } = useToken();
    const navigate = useNavigate();
    const handleSignOut = () => {
        clearToken();
        navigate("/", {replace: true})
        window.location.reload();
    }

    return (
        <div className="App">
            <nav className="navbar navbar-expand-lg navbar-light fixed-top">
                <div className="container">
                    <Link className="navbar-brand" to={"/sign-in"}>Calorie Tracker</Link>
                    <div className="collapse navbar-collapse" id="navbarTogglerDemo02">
                        {!token &&
                            <ul className="navbar-nav ml-auto">
                                <li className="nav-item">
                                    <Link className="nav-link" to={"/login"}>Login</Link>
                                </li>
                                <li className="nav-item">
                                    <Link className="nav-link" to={"/sign-up"}>Sign up</Link>
                                </li>
                            </ul>}
                        {token &&
                            <ul className="navbar-nav ml-auto">
                                <li className="nav-item">
                                    <Link className="nav-link" to={"/"} onClick={handleSignOut}>Sign out</Link>
                                </li>
                            </ul>
                        }
                    </div>
                </div>
            </nav>

            <div className="auth-wrapper">
                <div className="auth-inner">
                    <Routes>
                        <Route path="/" element={<Home/>}/>
                        <Route path="/login" element={<Login setToken={setToken}/>}/>
                        <Route path="/sign-up" element={<SignUp/>}/>
                        <Route path="/dashboard" element={<Dashboard setToken={setToken}/>}/>
                    </Routes>
                </div>
            </div>
        </div>
    );
}

export default App;