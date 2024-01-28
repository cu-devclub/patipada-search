import { CheckIcon, CloseIcon } from "@chakra-ui/icons";
import {
  useEditableControls,
  ButtonGroup,
  IconButton,
  Editable,
  Tooltip,
  EditablePreview,
  useColorModeValue,
  EditableInput,
  Textarea,
} from "@chakra-ui/react";

interface EditableInputFormProps {
  defaultValue: string;
}

function EditableInputForm({ defaultValue }: EditableInputFormProps) {
  /* Here's a custom control */
  function EditableControls() {
    const {
      isEditing,
      getSubmitButtonProps,
      getCancelButtonProps,
      // getEditButtonProps,
    } = useEditableControls();

    return isEditing ? (
      <ButtonGroup justifyContent="end" size="sm" w="full" spacing={2} mt={2}>
        <IconButton
          aria-label="confirm"
          icon={<CheckIcon />}
          {...getSubmitButtonProps()}
        />
        <IconButton
          aria-label="cancel"
          icon={<CloseIcon boxSize={3} />}
          {...getCancelButtonProps()}
        />
      </ButtonGroup>
    ) : null;
  }

  return (
    <Editable
      defaultValue={defaultValue}
      isPreviewFocusable={true}
      selectAllOnFocus={false}
      fontWeight={"light"}
      w="full"
      h="full"
    >
      <Tooltip label="Click to edit" shouldWrapChildren={true}>
        <EditablePreview
          _hover={{
            background: useColorModeValue("gray.100", "gray.700"),
          }}
          w="full"
          h="full"
        />
      </Tooltip>
      <EditableInput as={Textarea} fontWeight={"light"} w="full" h="full" />
      <EditableControls />
    </Editable>
  );
}
export default EditableInputForm;
