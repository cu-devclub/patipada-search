import wordCloud from "/word-cloud.png";
import { Image } from "@chakra-ui/react";

const WordCloud = () => {
  return <Image src={wordCloud} alt="WordCloud" objectFit="contain" />;
};

export default WordCloud;
