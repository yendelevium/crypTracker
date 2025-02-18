// isLoggedIn
// User Details
import {create} from 'zustand'
import { type TCoin, TUser } from '../utils/types'

interface UserState{
    isLoggedIn: boolean
    currentUser: null|TUser
    watchlist: TCoin[]
    toastMessage: string|null
    setIsLoggedIn: (loginState: boolean)=>void
    setCurrentUser: (newUser: TUser)=>void
    setWatchlist: (newWatchlist: TCoin[]) => void
    setToastMessage: (message: null|string)=>void
    handleUserLogin: (username: string, password:string) => any
    handleUserSignup: (username: string, password:string) => any
}

// When I close the tab, the cookie persists, but the storeData DOESN'T
// Do something abt this

// Check the cookie, if it's there, decrypt it and set userState
// On refresh itself it goes off TF
const userStore = create<UserState>((set)=>({
    isLoggedIn: false,
    currentUser: null,
    watchlist: [],
    toastMessage: null,
    setToastMessage: (message: string|null)=>{set(()=>({toastMessage: message}))},
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