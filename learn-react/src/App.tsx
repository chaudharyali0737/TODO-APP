import React from 'react';
import './App.css';
function App() {
  return (<>
    <img src="logo512.png" width="60px" height="60" alt="logo" />
    <nav>
      <h1>MesissaSoft Services</h1>
      <ul>
        <li>Services</li>
        <li>Pricing</li>
        <li>About</li>

      </ul>
    </nav>
  </>

  );
}
function React_benifits() {
  return (
    <>
      <div>
        <img src="/logo512.png" width="60px" height="60" alt="logo" />
        <nav>
          <h1>fun Facts About React</h1>
          <ul>
            <li>it was fist used in 2013</li>
            <li>was originally created by jordan walke</li>
            <li>has well over 100k stars on github</li>
            <li>is maintained by facebook</li>
            <li>empowers thousands of enterprise apps , including mobile apps</li>

          </ul>
        </nav>
      </div>

    </>
  );
}
export { App, React_benifits }
