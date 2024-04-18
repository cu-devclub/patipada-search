import { Button } from "@chakra-ui/react";
import BaseModal from "./BaseModal";

interface ActionModalProps {
  modalTitle?: string;
  modalBody: string | React.ReactNode;
  openModal: boolean;
  closeModal: () => void;
  confirm: () => void;
}

function ActionModal({
  modalTitle,
  modalBody,
  openModal,
  closeModal,
  confirm,
}: ActionModalProps) {
  return (
    <BaseModal
      modalTitle={modalTitle}
      modalBody={modalBody}
      openModal={openModal}
      closeModal={closeModal}
      ActionButtons={
        <>
          <Button variant="cancel" mr={3} onClick={closeModal}>
            ยกเลิก
          </Button>
          <Button variant="success" onClick={confirm}>
            ยืนยัน
          </Button>
        </>
      }
    />
  );
}

export default ActionModal;
