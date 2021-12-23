import React from "react";
import Login from "./login.component";

export default function Dashboard({setToken}) {

    if (!sessionStorage.getItem("token")) {
        return <Login setToken={setToken}/>
    }
    return(
        <form>
            <h3>Dashboard For the Calorie Tracke App</h3>
        </form>
    );
}