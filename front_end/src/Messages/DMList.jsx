import { React, useState } from "react";
import pPic from './pp.png';

function DMList(props) {
    return (
    <div class="flex flex-col min-h-screen">
        <div class="bg-gray-600 text-center text-white text-lg p-3">
            navBar
        </div>

        <main class="flex-1 overflow-y-auto flex flex-col">
            <div class="bg-purple-300 flex py-6 pl-20 pr-18 border-2 border-purple-700">
                <img src={pPic} class="pr-12 w-28" alt="a placeholder profile picture"></img>
                <div>
                    <div class="font-bold text-xl pb-4">rainethhh</div>
                    <div class="text-lg ml-10">
                        Yes I agree, it really does feel like that
                    </div>
                </div>
            </div>

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
        </main>

        <div class="bg-purple-700 text-center text-white text-xl p-3">
            Wormhole Chat
        </div>
    </div>
    );

}

export default DMList;
