import { Button } from "@chakra-ui/react";
import BaseModal from "./BaseModal";
import React from "react";

interface InformModalProps {
  modalTitle?: string;
  modalBody: string | React.ReactNode;
  modalButtonText?: string;
  openModal: boolean;
  closeModal: () => void;
}

function InformModal({
  modalTitle,
  modalBody,
  modalButtonText,
  openModal,
  closeModal,
}: InformModalProps) {
  return (
    <BaseModal
      modalTitle={modalTitle}
      modalBody={modalBody}
      openModal={openModal}
      closeModal={closeModal}
      ActionButtons={
        <Button colorScheme="orange" mr={3} onClick={closeModal} color="white">
          {modalButtonText}
        </Button>
      }
    />
  );
}

export default InformModal;
