import sideImg from "../assets/Purple-BG.jpg"

function Home(){
    return(
        <main className="p-3">
            <h1 className="text-4xl pb-2">Home/crypTracker</h1>
            <div>
                <p className="text-2xl">Welcome, to crypTracker : A real time cryptocurrency tracker to stay on top of your precious coins!</p>
                <div className="flex gap-2 p-1">
                    <div className="flex-1 p-3">
                        <p>
                            I'm Yash, and as of writing this, I'm a sophomore in college:) As you can see the way this website is looking, I'm not a good designer lol.
                            My favourite colour is purple, which again, you might have figured it out idk. 
                        </p>
                        <p className="my-1">
                            I made this website for myself, as all the other apps have SO MUCH GOING ON and I just get lost in the numbers lol.
                            So, I decided to go with a simple UI with just some of the most popular coins (because there's like 14,000+ coins) so you can know what's hot in the crypto world.
                        </p>
                        <p className="my-1">
                            I also learnt A LOT while making this site: TypeScript, websockets,  and JWTs.
                            I also tried out a new Go framework - Fiber, a new ORM in GORM, TailwindCSS, and Zustand. Apart from these, I also learnt a lot of stuff about Go and React as well.
                            And the best part? This was all my work. Didn't blindly follow a "create ur own crypto tracker" from youtube(if it even exists?), no copy-pasting code from ChatGPT : just raw-dogging the documentation and stackoverflow/reddit threads as God intended.
                            This was INSANELY FUN. I was up till 4-5am on countless nights, just coding and watching my idea come to life. Anyways, see ya! Use this app if you want to lol.
                            Here's my <a href="https://github.com/yendelevium" className="text-emerald-500">GitHub</a> if fancy seeing what I'm up to now hehe.
                        </p>
                    </div>
                    <div className="flex-1 p-3">
                        <img src={sideImg} alt="side-bg" className="rounded-3xl h-150 w-120 mx-auto object-cover"/>
                    </div>
                </div>

            </div>
        </main>
    )
}

export default Home