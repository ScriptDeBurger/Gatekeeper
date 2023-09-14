# GateKeeper

1	CONTEXT

Elk Fonteyn vakantiepark heeft een parkeergelegenheid. De directie van de parken vindt het belangrijk dat hun klanten altijd hun auto kunnen parkeren op de bijgelegen parkeerplaats. Het parkeerterrein is echter momenteel publiek toegankelijk. Dit zorgt ervoor dat ook mensen die geen accommodatie hebben gehuurd, gebruik maken van de parkeergelegenheid. Hierdoor kunnen klanten soms geen plek meer vinden op de parkeerplaats.

1.1	OPLOSSINGSRICHTING

De directie wil alle parkeerplaatsen gaan voorzien van een slagboom. Met een camera wordt het kenteken gescand. De klant kan bij het maken van een reservering van een accommodatie, het kenteken van de auto opgeven. Als het gescande kenteken wordt herkend, het kenteken is gekoppeld aan een reservering en de klant is ingecheckt, dient de slagboom open te gaan. Om te controleren of een klant is ingecheckt, controleert het systeem of de huidige datum en tijd ligt tussen de begin- en eindtijd van de gekoppelde reservering.

1.2	EISEN

De directie stelt de volgende eisen aan de applicatie:
Functionele eisen

- Een klant dient bij zijn reservering een kenteken te kunnen registreren zodat enkel één auto gekoppeld aan een reservering toegang heeft tot het parkeerterrein
- Het systeem dient a.d.h.v. een kenteken te bepalen of het voertuig toegang heeft tot het parkeerterrein
- Het kenteken dient te worden herkend
- Het kenteken dient te zijn gekoppeld aan een reservering
- Het tijdstip van de scan dient tussen begin- en einddatum van de reservering te liggen

Technische eisen
- De applicatie dient te worden geprogrammeerd in Golang
- De applicatie dient (in eerste instantie) een console app te zijn
- De code dient te worden opgeslagen in een versiebeheersysteem

1.3	AANPAK
De directie wil graag zo snel mogelijk een eerste versie van de applicatie zien maar snapt ook dat niet alles direct kan worden gerealiseerd. Vandaar dat met een eenvoudige applicatie wordt begonnen die iedere keer wordt uitgebreid. Zo kun je al snel iets laten zien en heb je altijd wel iets af. In de volgende hoofdstukken wordt beschreven wat er in de verschillende versies wordt verwacht van de “GateKeeper” app.
 
