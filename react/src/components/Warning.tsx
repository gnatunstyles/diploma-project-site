import { useState } from "react";
import styles from "../styles/warning.module.sass";

export default function Warning() {
  const [show, setShow] = useState(false);
  return (
    <div className={styles.container}>
      {show && <div className={styles.message}>
        Amount of points is above 10 million points. It can be displayed incorrectly. <br/>Please, make processing for better experience.
        </div>}
      <div className={styles.warning}
        onMouseEnter={() => {
          setShow(true);
        }}
        onMouseLeave={() => {
          setShow(false);
        }}
      >
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
          <defs>
            <linearGradient
              gradientTransform="matrix(1.31117 0 0 1.30239 737.39 159.91)"
              gradientUnits="userSpaceOnUse"
              id="0"
              y2="-.599"
              x2="0"
              y1="45.47"
            >
              <stop stop-color="#ffc515" />
              <stop offset="1" stop-color="#ffd55b" />
            </linearGradient>
          </defs>
          <g transform="matrix(.85714 0 0 .85714-627.02-130.8)">
            <path
              d="m797.94 212.01l-25.607-48c-.736-1.333-2.068-2.074-3.551-2.074-1.483 0-2.822.889-3.569 2.222l-25.417 48c-.598 1.185-.605 2.815.132 4 .737 1.185 1.921 1.778 3.404 1.778h51.02c1.483 0 2.821-.741 3.42-1.926.747-1.185.753-2.667.165-4"
              fill="url(#0)"
            />
            <path
              d="m-26.309 18.07c-1.18 0-2.135.968-2.135 2.129v12.82c0 1.176.948 2.129 2.135 2.129 1.183 0 2.135-.968 2.135-2.129v-12.82c0-1.176-.946-2.129-2.135-2.129zm0 21.348c-1.18 0-2.135.954-2.135 2.135 0 1.18.954 2.135 2.135 2.135 1.181 0 2.135-.954 2.135-2.135 0-1.18-.952-2.135-2.135-2.135z"
              transform="matrix(1.05196 0 0 1.05196 796.53 161.87)"
              fill="#000"
              stroke="#40330d"
              fill-opacity=".75"
            />
          </g>
        </svg>
      </div>
    </div>
  );
}
