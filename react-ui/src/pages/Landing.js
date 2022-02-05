import React, { useContext } from 'react';
import Header from '../sections/Header';
import { Navigate } from 'react-router-dom';
import { UserContext } from '../hooks/UserContext';

export default function Landing() {
    const { user } = useContext(UserContext);
        if(user) {
            return <Navigate to='/home' />
        }

    return(
        <div className="page">
            <Header/>
           <h2>Calorie Tracker</h2>
        </div>
    )
}