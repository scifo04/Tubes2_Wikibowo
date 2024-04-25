import React from 'react';
import './ResList.css';

const ResList = ({ tempLinks, urls, exec, len }) => {
  // Conditionally render based on whether the list is empty or not
  if (tempLinks.length === 0) {
    return <p></p>;
  }

  // Display list items if the list is not empty
  return (
    <div className='third'>
      <p>Racing done in <b>{exec}</b> ms after traversing through <b>{len}</b> articles</p>
      <p>Link visited: <b>{tempLinks.length} links</b></p>
      {tempLinks.map((item, index) => (
        <p>{index+1}. <a href={urls[index]} style={{textDecoration:"none",color:"white"}}>{item}</a></p>
      ))}
    </div>
  );
};

export default ResList