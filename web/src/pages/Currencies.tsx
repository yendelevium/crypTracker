import React from "react"
import coinStore from "../store/coinStore"
import Coin from "../components/Coin"
import Table from "../components/Table"
import Toast from "../components/Toast"
import toastStore from "../store/toastStore"

// Invalid DOM property `stroke-linecap`. Did you mean `strokeLinecap`? Component Stack: 
// This is an SVG error, where react needs the stuff to be camelCased
// The Error is from the Search-Icon SVQ in the search-bar

function Currencies(){
    const {allCoins, setAllCoins} = coinStore()
    const {toastMessage} = toastStore()
    const [search,setSearch] = React.useState<string>("")
    React.useEffect(()=>{
        // TODO:
        // Error Handling needed
        const getCoins = async ()=>{
            const response = await fetch("http://localhost:8080/crypto/coins")
            if (!response.ok) {
                console.log(response.status)
            }
            const coinJSON = await response.json()
            setAllCoins(coinJSON)
            // console.log(coinJSON)
        }
        getCoins()
    },[setAllCoins])

    function searchHandler(event: React.ChangeEvent<HTMLInputElement>){
        const {value}:{value: string, name: string} = event.currentTarget
        setSearch(value)
        console.log(value)
    }

    const coinJSX = allCoins.map(coin => {
        return(
            <Coin coinData={coin} key={coin.id} filter={search.trim()}/>
        )
    })

    return(
        <main className="p-3">
            <div className="flex justify-between">
                <h1 className="text-4xl pb-2">Currencies</h1>
                {/* Search Bar */}
                <div className="flex items-center justify-between flex-column flex-wrap md:flex-row space-y-4 md:space-y-0 pb-4 bg-white">
                    <label htmlFor="table-search" className="sr-only">Search</label>
                    <div className="relative">
                        <div className="absolute inset-y-0 rtl:inset-r-0 start-0 flex items-center ps-3 pointer-events-none">
                            <svg className="w-4 h-4 text-gray-500" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
                            </svg>
                        </div>
                    {/* TODO : Implement dynamic search-bar */}
                        <input type="text" id="table-search-users" className="block p-2 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg w-80 bg-gray-50 focus:ring-blue-500 focus:border-blue-500" placeholder="Search for coin(s)" value={search} onChange={(event)=>searchHandler(event)} name="search"/>
                    </div>
                </div>
            </div>

            <div className="mb-2">
                <p>Checkout all the available crypto-currencies! Login to add any coin to your watchlist.</p>
                <p></p>
            </div>

            {/* Coin display Table */}
            <Table coinJSX={coinJSX}/>
            {toastMessage ? <Toast /> : null}
        </main>
    )
}

export default Currencies