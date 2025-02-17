import { create } from 'zustand'
import { type TCoin } from '../utils/types'

interface CoinState {
    allCoins: TCoin[];
    setAllCoins: (newCoins: TCoin[]) => void;
}

const coinStore = create<CoinState>((set) => ({
    allCoins: [],
    setAllCoins: (newCoins) => set(() => ({ allCoins: newCoins })),
}));

export default coinStore;
