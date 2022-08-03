import React from "react";
import { useSpeechSynthesis } from "react-speech-kit";
const Speech = (comm) => {
    const [value, setValue] = React.useState("");
    const { speak,voices } = useSpeechSynthesis();
    // const {voices} = speechSynthesis.getVoices()
    // console.log(voices.length)
    // console.log(voices)
    // console.log("inside speech "+comm)
    // speak({text:comm})
    return (
        <div className="speech">
            <div className="group">
                <h2>Text To Speech Converter Using React Js</h2>
            </div>
            <div className="group">
        <textarea
            rows="10"
            value={value}
            onChange={(e) => setValue(e.target.value)}
        ></textarea>
            </div>
            <div className="group">
                <button onClick={() => {
                    console.log(voices)
                    speak({ text: value, voice: voices[0]})
                }}>
                    Speech
                </button>
            </div>
        </div>
    );
};
export default Speech;