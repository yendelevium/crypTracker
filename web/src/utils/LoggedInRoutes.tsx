import { Outlet, Navigate } from 'react-router'
import userStore from '../store/userStore'

function LoggedInRoutes() {
    const {currentUser} = userStore()
    return currentUser ? <Navigate to={"/cryptocurrencies"}/> : <Outlet /> 
}

export default LoggedInRoutes