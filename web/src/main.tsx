import { StrictMode } from 'react'
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

import ProtectedRoutes from './utils/ProtectedRoutes';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path='/' element={<Home />}/>

        <Route path='/login' element={<Login />}/>
        <Route path='/signup' element={<Signup />}/>
        <Route path='/cryptocurrencies' element={<Currencies />}/>

        <Route element={<ProtectedRoutes/>}>
          <Route path='/profile' element={<Profile />}/>
          <Route path='/watchlist' element={<Watchlist />}/>
        </Route>

      </Routes>
    </BrowserRouter>
  </StrictMode>
)
