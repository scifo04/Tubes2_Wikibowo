import React, { useState } from 'react'; 
import './App.css';
import LinkForm from './components/LinkForm';
import SearchButton from './components/SearchButton';
import EndForm from './components/EndForm';
import OnOff from './components/OnOff';
import "./fonts/Inter-Bold.ttf"

function Home() {
  const [tempLinks, setTempLinks] = useState('');
  const [linkValue, setLinkValue] = useState('');
  const [finValue, setFinValue] = useState('');
  const [isOn, setIsOn] = useState(false);

  console.log(tempLinks);

  return (
    <div>
        <h1 style={{fontFamily:"Inter"}}>WIKIBOWO DA WIKIRACER</h1>
        <div>
          <p>"Balapan adalah jalan hidupku. Curang pake Website adalah cara menangku"</p>
          <p>- Sun Tzu, Art of War</p>
        </div>
        <div>
          <OnOff isOn={isOn} setIsOn={setIsOn}/>
        </div>
        <div>
          <LinkForm linkValue={finValue} setLinkValue={setLinkValue}/>
          <EndForm finValue={finValue} setFinValue={setFinValue}/>
        </div>
        <SearchButton isOn={isOn} linkValue={linkValue} finValue={finValue} tempLinks={tempLinks} setTempLinks={setTempLinks}/>
    </div>
  );
}

export default Home;