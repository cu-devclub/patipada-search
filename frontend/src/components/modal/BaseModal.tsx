import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalCloseButton,
  ModalBody,
  ModalFooter,
} from "@chakra-ui/react";

interface BaseModalProps {
  modalTitle?: string;
  modalBody: React.ReactNode | string;
  openModal: boolean;
  closeModal: () => void;
  ActionButtons?: React.ReactNode;
}

function BaseModal({
  modalTitle,
  modalBody,
  openModal,
  closeModal,
  ActionButtons,
}: BaseModalProps) {
  return (
    <Modal isOpen={openModal} onClose={closeModal}>
      <ModalOverlay />
      <ModalContent mx={4}>
        <ModalHeader> {modalTitle} </ModalHeader>
        <ModalCloseButton />
        <ModalBody>{modalBody}</ModalBody>
        <ModalFooter>{ActionButtons}</ModalFooter>
      </ModalContent>
    </Modal>
  );
}

export default BaseModal;
