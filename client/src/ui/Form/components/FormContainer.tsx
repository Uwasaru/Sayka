import { FC } from "react";
import { FormProvider, useForm } from "react-hook-form";

type TProps = {
  submitFunction?: (data: any) => void;
  className?: string;
  id?: string;
  children: React.ReactNode;
};

export const FormContainer: FC<TProps> = ({
  submitFunction,
  className,
  id,
  children,
}) => {
  const methods = useForm();

  const onSubmit = (data: any) => {
    submitFunction?.(data);
  };

  return (
    <FormProvider {...methods}>
      <form
        onSubmit={methods.handleSubmit(onSubmit)}
        className={className}
        id={id}>
        {children}
      </form>
    </FormProvider>
  );
};
