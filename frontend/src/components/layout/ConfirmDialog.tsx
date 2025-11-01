import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";

interface ConfirmDialogProps {
  open: boolean;
  onConfirm: () => void;
  onCancel: () => void;
  title?: string;
  message?: string;
}

export const ConfirmDialog = ({
  open,
  onConfirm,
  onCancel,
  title = "Are you sure?",
  message = "This action cannot be undone.",
}: ConfirmDialogProps) => (
  <Dialog open={open} onOpenChange={onCancel}>
    <DialogContent>
      <button
        onClick={onCancel}
        className="absolute top-3 right-3 text-gray-500 hover:text-gray-700"
      ></button>
      <DialogHeader>
        <DialogTitle>{title}</DialogTitle>
      </DialogHeader>
      <p className="text-sm text-gray-600 dark:text-gray-300">{message}</p>
      <DialogFooter className="flex justify-end gap-2 mt-4">
        <Button variant="outline" onClick={onCancel}>
          No
        </Button>
        <Button variant="destructive" onClick={onConfirm}>
          Yes
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
);
