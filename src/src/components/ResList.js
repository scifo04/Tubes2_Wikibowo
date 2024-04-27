import React from 'react';
import './ResList.css';

const ResList = ({resultResponse }) => {
  // Conditionally render based on whether the list is empty or not
  if (resultResponse.resultLink.length === 0) {
    return <p></p>;
  }

  // Display list items if the list is not empty
  return (
    <div className='third'>
      <p>Racing done in <b>{resultResponse.exec}</b> ms after traversing through <b>{resultResponse.len}</b> articles</p>
      <p>Links Visited/Depth: <b>{resultResponse.resultLink.length}</b> links</p>
      {resultResponse.resultLink.map((item, index) => (
        <p key={index}>{index+1}. <a href={resultResponse.urls[index]} style={{textDecoration:"none",color:"white"}}>{item}</a></p>
      ))}
    </div>
  );
};

export default ResList