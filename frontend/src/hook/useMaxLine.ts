import { useEffect, useState } from 'react';

function useMaxLines() {
    const [maxLines, setMaxLines] = useState<number>(() => {
        // Define breakpoints and corresponding maxLines values
        const breakpoints = {
            sm: 480, // Example breakpoint for small screens
            md: 820, // Example breakpoint for medium screens
            lg: 1024, // Example breakpoint for large screens
            xl: 1280, // Example breakpoint for extra-large screens
        };

        // Determine the maxLines value based on screen width
        const screenWidth = window.innerWidth;

        if (screenWidth < breakpoints.sm) {
            return 4; // Adjust the maxLines value for small screens
        } else if (screenWidth < breakpoints.md) {
            return 6; // Adjust the maxLines value for medium screens
        } else {
            return 8; // Default maxLines value for larger screens
        }
    });

    useEffect(() => {
        const breakpoints = {
            sm: 480,
            md: 820,
            lg: 1024,
            xl: 1280,
        };

        const handleResize = () => {
            const screenWidth = window.innerWidth;

            if (screenWidth < breakpoints.sm) {
                setMaxLines(2);
            } else if (screenWidth < breakpoints.md) {
                setMaxLines(6);
            } else {
                setMaxLines(8);
            }
        };

        window.addEventListener('resize', handleResize);

        return () => {
            window.removeEventListener('resize', handleResize);
        };
    }, []);

    return maxLines;
}

export default useMaxLines;
