import { AspectRatio } from "@chakra-ui/react";
import { forwardRef, useImperativeHandle } from "react";
import { timeToSeconds } from "../../functions";
import React from "react";

interface VdoProps {
  question: string;
  startTime: string;
  endTime: string;
  youtubeURL: string;
}

export interface VdoRef {
  replay: () => void;
}

const YoutubeVideo = forwardRef<VdoRef, VdoProps>(
  ({ startTime, endTime, youtubeURL, question }: VdoProps, ref) => {
    const displayStartTime = timeToSeconds(startTime);
    const displayEndTime = timeToSeconds(endTime);
    const displayYoutubeURL = `https://www.youtube.com/embed/${youtubeURL}?start=${displayStartTime}&end=${displayEndTime}`;

    const replay = () => {
      const iframe = document.getElementById(question) as HTMLImageElement;
      if (iframe) {
        iframe.src = displayYoutubeURL;
      }
    };

    useImperativeHandle(ref, () => ({
      replay,
    }));

    return (
        <AspectRatio maxW={["560px"]} maxH="300px" ratio={1}>
          <iframe
            id={question}
            title={question}
            src={displayYoutubeURL}
            allowFullScreen
          />
        </AspectRatio>
    );
  }
);

export default YoutubeVideo;
