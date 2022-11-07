Details
========

Le projet est composer de plusieurs partie utile au bon fonctionnement du projet.

Authentification
-----------------

Nous avons mis en place un system d'Authentification sur notre projet.
Pour ce faire la personne doit renseigner uniquement un Username et un password associer.

C'est information sont stocker dans la base de donner et chaque utilisateur est associer à un token.


cela permet de personnaliser les action selon l'utilisteur.


.. _service:

Services
---------

:doc:`api`

.. _action:

Action 
-------

L'une des composantes les plus importante du projet sont les Actions.
cette dernière est une taches réaliser sur un service suivant des instruction déterminer en amont.

Voici un diagramme explicant de façons large le fonctionnement :

.. image:: images/action.png
    :width: 400


Pour ce faire une connection au service est necessaire pour le fonctionnement voir ci-joint pour plus d'information: :ref:`service`

quelque point clés pour comprendre le composant.

Tous d'abord le composant Action est lier a un trigger définie en amont pour guider l'action a faire.
sans un trigger il ne serai pas quoi faire.

Donc des trigger on été définie pour chaque action et reaction.
Un exemple de trigger :
    
    Connection au gmail
    
    check tous les x temps si nous avons reçus un mail
    
    stockage des information du mail.

.. _reaction:

Reaction
---------
