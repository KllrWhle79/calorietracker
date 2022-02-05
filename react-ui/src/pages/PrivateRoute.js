import React, {useContext} from 'react';
import {Navigate, useLocation} from 'react-router-dom';
import {UserContext} from '../hooks/UserContext';
import Loading from './../components/Loading';


export default function PrivateRoute({children}: { children: JSX.Element }) {
    const {user, isLoading} = useContext(UserContext);
    const location = useLocation();

    if (isLoading) {
        return <Loading/>
    }

    if (user) {
        return children;
    } else {
        return <Navigate to='/login' state={{from: location}}/>
    }

}




