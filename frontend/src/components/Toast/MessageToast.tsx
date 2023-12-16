import { ToastStatusType } from "../../models/toast";
import { useToast } from "@chakra-ui/react";

/**
 * Creates a message toast and returns a function to add new toasts.
 *
 * @param {MessageToastProps} newRes - The properties of the new toast.
 * @return {void}
 */
interface MessageToastProps {
  description: string;
  status: ToastStatusType;
}

const MessageToast = () =>  {
  const toast = useToast();
  // types are: "success", "info", "warning", "error"

  const addToast = (newRes: MessageToastProps) => {
    toast({
      description: newRes.description,
      status: newRes.status,
      position: "top",
      isClosable: true,
      duration: 5000,
      variant: "left-accent",
    });
  };

  return { addToast };
}

export default MessageToast;
