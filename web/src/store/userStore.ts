// isLoggedIn
// User Details
import {create} from 'zustand'
import { type TCoin, TUser } from '../utils/types'

interface UserState{
    currentUser: null|TUser
    watchlist: TCoin[]
    setCurrentUser: (newUser: null|TUser)=>void
    setWatchlist: (newWatchlist: TCoin[]) => void
    handleUserLogin: (username: string, password:string) => any
    handleUserSignup: (username: string, password:string) => any
}

// When I close the tab, the cookie persists, but the storeData DOESN'T
// Do something abt this

// Check the cookie, if it's there, decrypt it and set userState
// On refresh itself it goes off TF
const userStore = create<UserState>((set)=>({
    currentUser: null,
    watchlist: [],
    setWatchlist: (newWatchlist: TCoin[]) => {set(()=>({watchlist: newWatchlist}))},
    setCurrentUser: (newUser: null|TUser)=>{set(()=>({currentUser: newUser}))},
    handleUserLogin: async(username: string, password: string)=>{
        const loginResponse = await fetch("/users/login",{
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password
            })
        })

        const userData = await loginResponse.json()
        if (loginResponse.status !== 200){
            throw new Error(`Counldn't login: ${userData.message}`)
        }

        // console.log(userData)
        set(()=>({currentUser: userData.user_data, isLoggedIn: true}))
        return userData
    },

    handleUserSignup: async(username: string, password: string)=>{
        const signupResponse = await fetch("/users/",{
            method: "POST",
            body: JSON.stringify({
                username: username,
                password: password
            })
        })

        const userData = await signupResponse.json()
        if (signupResponse.status !== 201){
            throw new Error(`Couldn't signup ${userData.message}`)
        }

        // console.log(userData)
        set(()=>({currentUser: userData.user_data, isLoggedIn: true}))
        return userData
    }
}))

export default userStore