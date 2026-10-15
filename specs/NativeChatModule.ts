import { TurboModule, TurboModuleRegistry } from 'react-native';

export interface Spec extends TurboModule {
  readonly getHello: (input: string) => string;
}

export default TurboModuleRegistry.getEnforcing<Spec>(
  'NativeChatModule',
);
