import { Link } from "react-router"
import React from "react";
import coinStore from "../store/coinStore";
import userStore from "../store/userStore"
import { jwtVerify } from "jose";
import Cookie from "js-cookie"

import { type TCoin } from "../utils/types";

function Navbar(){
    const {setAllCoins} = coinStore()

    // Shift these useEffects to someother place, they should be executed upon DOM Rendering
    React.useEffect(()=>{
        // Using WebSocket instead of Socket.io as my backend ka main.go is using
        // websockets in app.Use(). So i either had to change that or this and I decided to change this
        console.log("WS connection? Frm Navbar lmao")
        const socket = new WebSocket("ws://localhost:8080/ws");

        socket.onopen = () => {
            console.log("Connected to WebSocket server");
            socket.send("Hello from client!");
        };

        // Converting the event payload into JSON, and updating all the coins
        socket.onmessage = (event) => {
            let coinData: TCoin[] = JSON.parse(event.data)
            setAllCoins(coinData)
        };

        socket.onerror = (error) => {
            console.error("WebSocket error:", error);
        };

        socket.onclose = () => {
            console.log("WebSocket closed");
        };
      
    },[])

    // To simulate login-persistance, so I can stay logged in even AFTER i refresh/reload/exit out of the page
    const { currentUser, setIsLoggedIn, setCurrentUser }=userStore()
    React.useEffect(()=>{
        // Get the JWT from the cookie
        const AuthJWT = Cookie.get("Authorization")
        console.log("Cookie",AuthJWT)
        if(AuthJWT!=undefined){
            const getUser = async ()=>{
                // Verify JWT
                const { payload } = await jwtVerify(AuthJWT, new TextEncoder().encode("Demon King of Salvation"));
                console.log(payload);

                // Fetch userDetails
                const reponse = await fetch(`/users/${payload.user_id}`)
                const userData = await reponse.json()
                console.log(userData)
                setCurrentUser(userData)
                setIsLoggedIn(true)
            }
    
            getUser()
        }
        console.log(currentUser?.username)    
    },[])


    return(
        <>
        <nav className="flex justify-between p-3 bg-violet-400 text-xl text-white">
            <div>
                <Link to="/">crypTracker</Link>
            </div>
            <div className="navLinks flex justify-evenly">
                <Link to="cryptocurrencies">Currencies</Link>
                {/* Watchlist must redirect to login if  you're not logged in*/}
                <Link to="watchlist/<userId>">Watchlist</Link>
                {/* Either Login OR Profile must be displayed based on whether user is logged in or not */}
                <Link to="login">Login</Link>
                <Link to="signup">Signup</Link>
                <Link to="profile">Profile</Link>
            </div>
        </nav>
        </>
    )
}

export default Navbar