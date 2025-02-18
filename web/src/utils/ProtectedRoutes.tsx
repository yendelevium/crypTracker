import { Outlet, Navigate } from 'react-router'
import userStore from '../store/userStore'

function ProtectedRoutes() {
    const {currentUser} = userStore()
    return currentUser ? <Outlet /> : <Navigate to={"/login"}/>
}

export default ProtectedRoutes