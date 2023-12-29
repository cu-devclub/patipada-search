import { Box } from "@chakra-ui/react";
import background from "/AnimatedBackground.svg";

interface AnimatedBackgroundProps {
  children: React.ReactNode;
}

function AnimatedBackground({ children }: AnimatedBackgroundProps) {
  return (
    <Box
      position="absolute"
      w="100%"
      minH="100svh"
      bgImage={`url(${background})`}
      bgSize="cover"
      bgPos="center"
      zIndex="-1"
    >
        {children}
    </Box>
  );
}

export default AnimatedBackground;
