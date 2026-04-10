export const Logo = () => {
  return (
    <svg
      width="100%"
      height="100%"
      xmlns="http://www.w3.org/2000/svg"
      xmlnsXlink="http://www.w3.org/1999/xlink"
      viewBox="0 0 1000 1000"
      aria-hidden="true"
    >
      <defs>
        <clipPath id="clippath">
          <path
            fill="none"
            d="M785.59,385.24c0,87.16-70.64,157.8-157.79,157.8-25.62,0-49.8-6.11-71.18-16.94,5.57.85,11.26,1.3,17.06,1.3,62.3,0,112.81-50.51,112.81-112.81s-50.51-112.81-112.81-112.81c-45.68,0-85,27.18-102.74,66.24,8.58-79.05,75.53-140.57,156.86-140.57,87.15,0,157.79,70.64,157.79,157.79ZM925.88,367.16c.36,5.98.56,12.01.56,18.08,0,164.93-133.7,298.65-298.63,298.65s-298.65-133.72-298.65-298.65S462.87,86.6,627.8,86.6c16.92,0,33.5,1.43,49.65,4.13-54.39-23.61-114.38-36.73-177.46-36.73-246.32,0-446,199.68-446,446s199.68,446,446,446,446-199.68,446-446c0-46.27-7.05-90.89-20.12-132.84Z"
          />
        </clipPath>
        <linearGradient
          id="gradation"
          x1="119.58"
          y1="158.32"
          x2="863.55"
          y2="826.53"
          gradientUnits="userSpaceOnUse"
        >
          <stop offset="0" stopColor="#2bc4cf" />
          <stop offset=".52" stopColor="#572bcf" />
          <stop offset="1" stopColor="#cf2ba1" />
        </linearGradient>
      </defs>
      <g className="cls-3" clipPath="url(#clippath)">
        <rect
          className="cls-1"
          width="1000"
          height="1000"
          fill="url(#gradation)"
        />
      </g>
    </svg>
  );
};
