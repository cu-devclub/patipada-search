import { AspectRatio } from "@chakra-ui/react";
import { forwardRef, useImperativeHandle } from "react";
import { timeToSeconds } from "../../functions";

interface VdoProps {
  id: string;
  startTime: string;
  endTime: string;
  youtubeURL: string;
}

export interface VdoRef {
  replay: () => void;
}

const YoutubeVideo = forwardRef<VdoRef, VdoProps>(
  ({ startTime, endTime, youtubeURL, id }: VdoProps, ref) => {
    const displayStartTime = timeToSeconds(startTime);
    const displayEndTime = timeToSeconds(endTime);
    const displayYoutubeURL = `https://www.youtube.com/embed/${youtubeURL}?start=${displayStartTime}&end=${displayEndTime}`;

    const replay = () => {
      const iframe = document.getElementById(id) as HTMLImageElement;
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
          id={id}
          title={id}
          src={displayYoutubeURL}
          allowFullScreen
        />
      </AspectRatio>
    );
  }
);

export default YoutubeVideo;
