// isLoggedIn
// User Details
import {create} from 'zustand'
import { type TCoin, TUser } from '../utils/types'

interface UserState{
    isLoggedIn: boolean
    currentUser: null|TUser
    watchlist: TCoin[]
    setIsLoggedIn: (loginState: boolean)=>void
    setCurrentUser: (newUser: TUser)=>void
    setWatchlist: (newWatchlist: TCoin[]) => void
    handleUserLogin: (username: string, password:string)=>void
    handleUserSignup: (username: string, password:string)=>void
}

// When I close the tab, the cookie persists, but the storeData DOESN'T
// Do something abt this

// Check the cookie, if it's there, decrypt it and set userState
// On refresh itself it goes off TF
const userStore = create<UserState>((set)=>({
    isLoggedIn: false,
    currentUser: null,
    watchlist: [],
    setWatchlist: (newWatchlist: TCoin[]) => {set(()=>({watchlist: newWatchlist}))},
    setIsLoggedIn: (loginState: boolean)=>{set(()=>({isLoggedIn:loginState}))},
    setCurrentUser: (newUser: TUser)=>{set(()=>({currentUser: newUser}))},
    handleUserLogin: async(username: string, password: string)=>{
        const loginResponse = await fetch("/users/login",{
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password
            })
        })

        if (loginResponse.status !== 200){
            console.log("Send an ERROR Toast later")
            return
        }

        const userData = await loginResponse.json()
        console.log(userData)
        set(()=>({currentUser: userData.user_data, isLoggedIn: true}))
    },

    handleUserSignup: async(username: string, password: string)=>{
        const signupResponse = await fetch("/users/",{
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password
            })
        })

        if (signupResponse.status !== 201){
            console.log("Send an ERROR Toast later")
            return
        }

        const userData = await signupResponse.json()
        console.log(userData)
        set(()=>({currentUser: userData.user_data, isLoggedIn: true}))
    }
}))

export default userStore