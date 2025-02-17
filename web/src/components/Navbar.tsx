import { Link } from "react-router"
import React from "react";
import coinStore from "../store/coinStore";
import { type TCoin } from "../utils/types";

function Navbar(){
    const {setAllCoins} = coinStore()
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