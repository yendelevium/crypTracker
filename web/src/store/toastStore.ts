import {create} from "zustand"


interface ToastState {
    toastMessage: string|null
    toastType: "success"|"error"
    setToastMessage: (message: null|string)=>void
    setToastType : (type: "success"|"error") => void
}

const toastStore = create<ToastState>((set)=>({
    toastMessage: null,
    toastType: "success",
    setToastMessage: (message: string|null)=>{set(()=>({toastMessage: message}))},
    setToastType: (type: "success"|"error")=>{set(()=>({toastType: type}))},

}))

export default toastStore