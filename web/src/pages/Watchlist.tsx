import userStore from "../store/userStore"
import Coin from "../components/Coin"
import React from "react"
import Table from "../components/Table"
import Toast from "../components/Toast"
import toastStore from "../store/toastStore"

function Watchlist(){
    const {watchlist} = userStore()
    const {toastMessage} = toastStore()
    let watchlistJSX: React.JSX.Element[] = [];
    // Only on a non-empty list (otherwise we get an error), map over the watchlist to create the JSX
    if (watchlist.length > 0){
        watchlistJSX = watchlist.map(coin=>{
            return (<Coin coinData={coin} key={coin.id}/>)
        })
    }

    return(
        <main className="p-3">
            <h1 className="text-4xl pb-2">Watchlist</h1>
            <div className="mb-2">
                The cryptocurrencies you are interested in
            </div>

            {/* Remove the 'Add to Watchlist' Thingy */}
            <Table coinJSX={watchlistJSX}/>
            {toastMessage ? <Toast/> : null}
        </main>
    )
}

export default Watchlist