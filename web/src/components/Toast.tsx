import React from "react"
import toastStore from "../store/toastStore"

function Toast(){
    const ref = React.useRef<HTMLDivElement | null>(null)

    // Resetting toast
    // We are using useEffect as otheriwse re - rendering toast somehow re-renders navbar, and that is a React Error
    // This way it doesn't rerender navbar and no error
    const { setToastMessage, toastMessage, toastType } = toastStore()
    React.useEffect(()=>{
        setTimeout(() => {
            setToastMessage(null); // Clear the message after 5 secs
        }, 5000);
    },[toastMessage])

    function dismissToast(ref: React.RefObject<HTMLDivElement | null>){
        // Removing the toast by putting display:none
        ref.current?.classList.add("hidden")
        setToastMessage(null)
    }

    let toastJSX: React.JSX.Element;
    if (toastType==="success"){
        toastJSX = 
            (<div ref={ref} className="bg-emerald-100 p-3 flex gap-2 items-center border border-teal-900 rounded-sm">
                    <span>{toastMessage}</span>
                    <button onClick={()=>dismissToast(ref)} className="p-1 cursor-pointer">X</button>
            </div>)
    }else{
        toastJSX = 
            (<div ref={ref} className="bg-red-100 p-3 flex gap-2 items-center border border-red-900 rounded-sm">
                <span>{toastMessage}</span>
                <button onClick={()=>dismissToast(ref)} className="p-1 cursor-pointer">X</button>
            </div>)
    }
    return(
        <div className="toast flex justify-end">
            {toastJSX}
        </div>
    )
}

export default Toast