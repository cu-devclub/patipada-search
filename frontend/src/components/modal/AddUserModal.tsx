import BaseModal from "./BaseModal";

interface AddUserModalProps {
  modalTitle?: string;
  modalBody: JSX.Element | string;
  openModal: boolean;
  closeModal: () => void;
}

function AddUserModal({
  modalTitle,
  modalBody,
  openModal,
  closeModal,
}: AddUserModalProps) {
  return (
    <BaseModal
      modalTitle={modalTitle}
      modalBody={modalBody}
      openModal={openModal}
      closeModal={closeModal}
    />
  );
}

export default AddUserModal;
