import { StrictMode, useEffect } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from "react-router";
import Home from './pages/Home';
import Navbar from './components/Navbar';
import Currencies from './pages/Currencies';
import Profile from './pages/Profile';
import Login from './pages/Login';
import Watchlist from './pages/Watchlist';
import Signup from './pages/Signup';
import "./index.css"
import { io } from 'socket.io-client';

// Doing socket stuff in useEffect?
useEffect(()=>{
  const socket = io("http://localhost:8080");
  socket.on("connect", () => {
    console.log(socket.id);
  });

},[])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path='/' element={<Home />}/>

        <Route path='/login' element={<Login />}/>
        <Route path='/signup' element={<Signup />}/>
        <Route path='/profile' element={<Profile />}/>

        <Route path='/cryptocurrencies' element={<Currencies />}/>
        <Route path='/watchlist/:userId' element={<Watchlist />}/>
      </Routes>
    </BrowserRouter>
  </StrictMode>
)
