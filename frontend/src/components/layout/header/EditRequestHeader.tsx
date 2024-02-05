import { Heading } from "@chakra-ui/react";
import BaseHeader from "./BaseHeader";

interface EditRequestHeaderProps {
  activePage: string;
  requestID?: string;
}
function EditRequestHeader({ activePage, requestID }: EditRequestHeaderProps) {
  return (
    <BaseHeader>
        <Heading>
          {requestID
            ? "ยื่นคำขอแก้ไขเนื้อหา คำขอเลขที่ " + requestID
            : activePage}
        </Heading>
    </BaseHeader>
  );
}

export default EditRequestHeader;
