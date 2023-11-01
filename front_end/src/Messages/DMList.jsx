import { React, useState,useEffect } from "react";
import pPic from './pp.png';

function DMList(props) {

    const examplefriends = [
        {username:"rainethhh", 
        msg:"Yes I agree, it really does feel like that"},
        {username:"rainethhh", 
        msg:"Yes I agree, it really does feel like that"},
        {username:"rainethhh", 
        msg:"Yes I agree, it really does feel like that"},
        {username:"rainethhh", 
        msg:"Yes I agree, it really does feel like that"},
        {username:"rainethhh", 
        msg:"Yes I agree, it really does feel like that"},
        {username:"rainethhh", 
        msg:"Yes I agree, it really does feel like that"} 
        ]
    
    const [msgs, setMsgs] = useState(null)
        useEffect(()=>{
            // ask back end for dms
            setMsgs(examplefriends)
        },[])

    return (
    <div class="flex flex-col min-h-screen">
        <div class="bg-gray-600 text-center text-white text-lg p-3">
            navBar
        </div>

        <div class="flex-1 overflow-y-auto flex flex-col">
            //only displays when connected with the backend
            {msgs ? msgs.map((msgObject, index) => (
                <div key = {index} class="bg-purple-300 flex py-6 pl-20 pr-18 border-2 border-purple-700">
                    <img src={pPic} class="pr-12 w-28" alt="a placeholder profile picture"></img>
                    <div>
                        <div class="font-bold text-xl pb-4">{msgObject.username}</div>
                        <div class="text-lg ml-10">
                            {msgObject.msg}
                        </div>
                    </div>
                </div> 
            )) : null}
            

            <div class="bg-purple-300 flex py-6 pl-20 pr-18 border-2 border-purple-700">
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
            </div>
        </div>

        <div class="bg-purple-700 text-center text-white text-xl p-3">
            Wormhole Chat
        </div>
    </div>
    );

}

export default DMList;
