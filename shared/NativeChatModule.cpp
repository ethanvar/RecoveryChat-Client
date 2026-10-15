#include "NativeChatModule.h"

namespace facebook::react {

NativeChatModule::NativeChatModule(std::shared_ptr<CallInvoker> jsInvoker)
    : NativeChatModuleCxxSpec(std::move(jsInvoker)) {}

std::string NativeChatModule::reverseString(jsi::Runtime &rt,
                                            std::string input) {
  return std::string(input.rbegin(), input.rend());
}

} // namespace facebook::react
