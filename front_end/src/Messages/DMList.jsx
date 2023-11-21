import { React, useState,useEffect } from "react";
import static1 from "../Static.js";
import pPic from '../images/pp.png';
import Background from '../Background/background'

function DMList(props) {
    const toggleHomepage = props.toggleHomepage
    const toggleDMMessage = props.toggleDMMessage

    const examplefriends = [
        {username:"rainethhh", 
        msg:"Yes I agree, it really does feel like that"},
        {username:"kingJames", 
        msg:"Appreciate it!!!"},
        {username:"Gene", 
        msg:"See ya"},
        {username:"kevonosdiaz", 
        msg:"Go Go Go Go Go"},
        {username:"duppy", 
        msg:"just grabbing a bakechef, be rigth there"},
        {username:"vicGPT", 
        msg:"The most optimal algorithm would be O(n log n)"} 
        ]
    
    const [msgs, setMsgs] = useState(null)
        useEffect(()=>{
            // ask back end for dms
            setMsgs(examplefriends)
        },[])

    return (
    <div className="flex flex-col items-center min-h-screen">
        <div className="bg-neutral-700 bg-opacity-90 text-black text-xl m-14 md:p-6 p-2 rounded-xl w-2/3 min-w-fit">
            <div className="flex-1 overflow-y-auto flex flex-col">
                {msgs ? msgs.map((msgObject, index) => (
                    <div  onClick = {() => {toggleDMMessage(msgObject.username)}} key = {index} className="bg-purple-200 hover:bg-purple-300 flex py-4 md:py-6 px-6 md:px-12 border-2 m-2 border-purple-700 rounded-xl">
                        <img src={pPic} className="pr-12 w-28 object-contain h-auto" alt="a placeholder profile picture"></img>
                        <div>
                            <div className="font-bold text-xl pb-4">{msgObject.username}</div>
                            <div className="text-lg ml-10 truncate">
                                {msgObject.msg}
                            </div>
                        </div>
                    </div> 
                )) : null}
            </div>
        </div>


            

            {/* <div class="bg-purple-300 flex py-6 pl-20 pr-18 border-2 border-purple-700">
                <img src={pPic} class="pr-12 w-28" alt="a placeholder profile picture"></img>
                <div>
                    <div class="font-bold text-xl pb-4">kingJames</div>
                    <div class="text-lg ml-10">
                        Appreciate it!!!
                    </div>
                </div>
            </div>

            <div class="bg-purple-300 flex py-6 pl-20 pr-18 border-2 border-purple-700">
                <img src={pPic} class="pr-12 w-28" alt="a placeholder profile picture"></img>
                <div>
                    <div class="font-bold text-xl pb-4">Gene</div>
                    <div class="text-lg ml-10">
                        See ya
                    </div>
                </div>
            </div>

            <div class="bg-purple-300 flex py-6 pl-20 pr-18 border-2 border-purple-700">
                <img src={pPic} class="pr-12 w-28" alt="a placeholder profile picture"></img>
                <div>
                    <div class="font-bold text-xl pb-4">kevonosdiaz</div>
                    <div class="text-lg ml-10">
                        Go Go Go Go Go
                    </div>
                </div>
            </div>

            <div class="bg-purple-300 flex py-6 pl-20 pr-18 border-2 border-purple-700">
                <img src={pPic} class="pr-12 w-28" alt="a placeholder profile picture"></img>
                <div>
                    <div class="font-bold text-xl pb-4">duppy</div>
                    <div class="text-lg ml-10">
                        just grabbing a bakechef, be right there
                    </div>
                </div>
            </div>

            <div class="bg-purple-300 flex py-6 pl-20 pr-18 border-2 border-purple-700">
                <img src={pPic} class="pr-12 w-28" alt="a placeholder profile picture"></img>
                <div>
                    <div class="font-bold text-xl pb-4">vicGPT</div>
                    <div class="text-lg ml-10">
                        The most optimal algorithm would be O(n log n)
                    </div>
                </div>
            </div> */}

        <div className="bg-purple-700 w-full text-center text-white text-xl p-3">
            Wormhole Chat
        </div>
    </div>
    );

}

export default DMList
