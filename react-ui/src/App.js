import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import {UserContext} from './hooks/UserContext';
import PrivateRoute from './pages/PrivateRoute';
import Register from './pages/Register';
import Login from './pages/Login';
import Landing from './pages/Landing';
import Home from './pages/Home';
import NotFound from './pages/NotFound';
import useFindUser from './hooks/useFindUser';
// import 'bootstrap/dist/css/bootstrap.min.css';

function App() {

    const {
        user,
        setUser,
        isLoading
    } = useFindUser();

    return (
        <Router>
            <UserContext.Provider value={{user, setUser, isLoading}}>
                <Routes>
                    <Route exact path="/" element={<Landing/>}/>
                    <Route path="/register" element={<Register/>}/>
                    <Route path="/login" element={<Login/>}/>
                    <Route exact path="/home" element={
                        <PrivateRoute>
                            <Home/>
                        </PrivateRoute>}>
                    </Route>
                    <Route element={<NotFound/>}/>
                </Routes>
            </UserContext.Provider>
        </Router>
    );
}

export default App;
