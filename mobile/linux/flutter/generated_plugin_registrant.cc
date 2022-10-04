//
//  Generated file. Do not edit.
//

// clang-format off

#include "generated_plugin_registrant.h"

<<<<<<< HEAD
<<<<<<< HEAD

void fl_register_plugins(FlPluginRegistry* registry) {
=======
=======
>>>>>>> 81baccd (feat(mobile): add flutter authentication system)
#include <url_launcher_linux/url_launcher_plugin.h>

void fl_register_plugins(FlPluginRegistry* registry) {
  g_autoptr(FlPluginRegistrar) url_launcher_linux_registrar =
      fl_plugin_registry_get_registrar_for_plugin(registry, "UrlLauncherPlugin");
  url_launcher_plugin_register_with_registrar(url_launcher_linux_registrar);
<<<<<<< HEAD
>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))
=======
=======

void fl_register_plugins(FlPluginRegistry* registry) {
>>>>>>> b07b2e5 (feat(mobile): add flutter authentication system)
>>>>>>> 81baccd (feat(mobile): add flutter authentication system)
}
