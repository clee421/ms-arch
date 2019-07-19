import { useState } from 'react';

export interface FormInput {
  value: string;
  onChange: (e: any) => void;
}

export interface FormInputConfig {
  type?: string;
  name?: string;
}

export function useFormInput(
  initialValue: string,
  config: FormInputConfig = {},
): FormInput {
  const [value, setValue] = useState(initialValue);

  function handleChange(e: any): void {
    setValue(e.target.value);
  }

  return {
    value,
    onChange: handleChange,
    ...config,
  };
}
