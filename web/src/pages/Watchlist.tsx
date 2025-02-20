import userStore from "../store/userStore"
import Coin from "../components/Coin"
import React from "react"
import Table from "../components/Table"
import Toast from "../components/Toast"
import toastStore from "../store/toastStore"

function Watchlist(){
    const {watchlist} = userStore()
    const {toastMessage} = toastStore()
    const [search,setSearch] = React.useState<string>("")
    let watchlistJSX: React.JSX.Element[] = [];
    // Only on a non-empty list (otherwise we get an error), map over the watchlist to create the JSX
    if (watchlist.length > 0){
        watchlistJSX = watchlist.map(coin=>{
            return (<Coin coinData={coin} key={coin.id} filter={search.trim()}/>)
        })
    }

    function searchHandler(event: React.ChangeEvent<HTMLInputElement>){
            const {value}:{value: string, name: string} = event.currentTarget
            setSearch(value)
            console.log(value)
    }

    return(
        <main className="p-3">
            <div className="flex justify-between">
                <h1 className="text-4xl pb-2">Watchlist</h1>
                {/* Search Bar */}
                <div className="flex items-center justify-between flex-column flex-wrap md:flex-row space-y-4 md:space-y-0 pb-4 bg-white">
                    <label htmlFor="table-search" className="sr-only">Search</label>
                    <div className="relative">
                        <div className="absolute inset-y-0 rtl:inset-r-0 start-0 flex items-center ps-3 pointer-events-none">
                            <svg className="w-4 h-4 text-gray-500" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
                            </svg>
                        </div>
                        <input type="text" id="table-search-users" className="block p-2 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg w-80 bg-gray-50 focus:ring-blue-500 focus:border-blue-500" placeholder="Search for coin(s)" value={search} onChange={(event)=>searchHandler(event)} name="search"/>
                    </div>
                </div>
            </div>
            <div className="mb-2">
                The cryptocurrencies you're interested in. Do your own research before buying anything because crypto is RISKY. But who am I to tell you what to do, you do you bestie.
            </div>

            {/* Remove the 'Add to Watchlist' Thingy */}
            <Table coinJSX={watchlistJSX}/>
            {toastMessage ? <Toast/> : null}
        </main>
    )
}

export default Watchlist