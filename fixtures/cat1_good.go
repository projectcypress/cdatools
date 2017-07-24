package fixtures

var Cat1Good = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="cda.xsl"?>
<ClinicalDocument xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="urn:hl7-org:v3" xmlns:voc="urn:hl7-org:v3/voc" xmlns:sdtc="urn:hl7-org:sdtc">
  <!-- QRDA Header -->
  <realmCode code="US"/>
  <typeId root="2.16.840.1.113883.1.3" extension="POCD_HD000040"/>
  <!-- US Realm Header Template Id -->
  <templateId root="2.16.840.1.113883.10.20.22.1.1"/>
  <!-- QRDA templateId -->
  <templateId root="2.16.840.1.113883.10.20.24.1.1"/>
  <!-- QDM-based QRDA templateId -->
  <templateId root="2.16.840.1.113883.10.20.24.1.2"/>
  <!-- This is the globally unique identifier for this QRDA document -->
  <id root="5b010313-eff2-432c-9909-6193d8416fac"/>
  <!-- QRDA document type code -->
  <code code="55182-0" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" displayName="Quality Measure Report"/>
  <title>QRDA Incidence Report</title>
  <!-- This is the document creation time -->
  <effectiveTime value="20130418154515"/>
  <confidentialityCode code="N" codeSystem="2.16.840.1.113883.5.25"/>
  <languageCode code="eng"/>
  <!-- reported patient -->
  <recordTarget>
    <patientRole>
      <!-- id root="Cypress" extension="515f07293054cfa39b00007b"/ -->
      <!-- Fake Medicare HIC number -->
      <id extension="12345" root="2.16.840.1.113883.4.572"/>
      <addr use="HP">
        <streetAddressLine>202 Burlington Rd.</streetAddressLine>
        <city>Bedford</city>
        <state>MA</state>
        <postalCode>01730</postalCode>
        <country>US</country>
      </addr>
      <telecom use="WP" value="tel:+1-781-271-3000"/>
      <patient>
        <name>
          <given>Norman</given>
          <family>Flores</family>
        </name>
        <administrativeGenderCode code="M" codeSystem="2.16.840.1.113883.5.1" codeSystemName="HL7 AdministrativeGender"/>
        <birthTime value="19890101083000"/>
        <raceCode code="1002-5" displayName="American Indian or Alaska Native" codeSystemName="CDC Race and Ethnicity" codeSystem="2.16.840.1.113883.6.238"/>
        <ethnicGroupCode code="2186-5" displayName="Not Hispanic or Latino" codeSystemName="CDC Race and Ethnicity" codeSystem="2.16.840.1.113883.6.238"/>
        <languageCommunication>
          <templateId root="2.16.840.1.113883.3.88.11.83.2" assigningAuthorityName="HITSP/C83"/>
          <templateId root="1.3.6.1.4.1.19376.1.5.3.1.2.1" assigningAuthorityName="IHE/PCC"/>
          <languageCode code="eng"/>
        </languageCommunication>
      </patient>
    </patientRole>
  </recordTarget>
  <!-- Example of an author who is a device -->
  <author>
    <time value="20130418154515"/>
    <assignedAuthor>
      <!-- id extension="Cypress" root="2.16.840.1.113883.19.5"/ -->
      <!-- NPI -->
      <id extension="FakeNPI" root="2.16.840.1.113883.4.6"/>
      <addr>
        <streetAddressLine>202 Burlington Rd.</streetAddressLine>
        <city>Bedford</city>
        <state>MA</state>
        <postalCode>01730</postalCode>
        <country>US</country>
      </addr>
      <telecom use="WP" value="tel:(781)271-3000"/>
      <assignedAuthoringDevice>
        <manufacturerModelName>Cypress</manufacturerModelName >
        <softwareName>Cypress</softwareName >
      </assignedAuthoringDevice >
    </assignedAuthor>
  </author>
  <custodian>
    <assignedCustodian>
      <representedCustodianOrganization>
        <id root="2.16.840.1.113883.19.5"/>
        <name>Cypress Test Deck</name>
        <telecom use="WP" value="tel:(781)271-3000"/>
        <addr>
          <streetAddressLine>202 Burlington Rd.</streetAddressLine>
          <city>Bedford</city>
          <state>MA</state>
          <postalCode>01730</postalCode>
          <country>US</country>
        </addr>
      </representedCustodianOrganization>
    </assignedCustodian>
  </custodian>
  <legalAuthenticator>
    <time value="20130418154515"/>
    <signatureCode code="S"/>
    <assignedEntity>
      <id root="bc01a5d1-3a34-4286-82cc-43eb04c972a7"/>
      <addr>
        <streetAddressLine>202 Burlington Rd.</streetAddressLine>
        <city>Bedford</city>
        <state>MA</state>
        <postalCode>01730</postalCode>
        <country>US</country>
      </addr>
      <telecom use="WP" value="tel:(781)271-3000"/>
      <assignedPerson>
        <name>
          <given>Henry</given>
          <family>Seven</family>
        </name>
      </assignedPerson>
      <representedOrganization>
        <id root="2.16.840.1.113883.19.5"/>
        <name>Cypress</name>
      </representedOrganization>
    </assignedEntity>
  </legalAuthenticator>
  <!-- TODO: This is where the provider information will go.       It is currently hard coded, but should be replaced with the providers       and the time over which they are performing. -->
  <documentationOf typeCode="DOC">
    <serviceEvent classCode="PCPR">
      <!-- care provision -->
      <effectiveTime>
        <low value="20100601"/>
        <high value="20100915"/>
      </effectiveTime>
      <!-- You can include multiple performers, each with an NPI, TIN, CCN. -->
      <performer typeCode="PRF">
        <time>
          <low value="20020716"/>
          <high value="20070915"/>
        </time>
        <assignedEntity>
          <!-- This is the provider NPI -->
          <id root="2.16.840.1.113883.4.6" extension="111111111" />
          <representedOrganization>
            <!-- This is the organization TIN -->
            <id root="2.16.840.1.113883.4.2" extension="1234567" />
            <!-- This is the organization CCN -->
            <id root="2.16.840.1.113883.4.336" extension="54321" />
          </representedOrganization>
        </assignedEntity>
      </performer>
    </serviceEvent>
  </documentationOf>
  <component>
    <structuredBody>
      <component>
        <section>
          <!--            *****************************************************************            Measure Section            *****************************************************************          -->
          <!-- This is the templateId for Measure Section -->
          <templateId root="2.16.840.1.113883.10.20.24.2.2"/>
          <!-- This is the templateId for Measure Section QDM -->
          <templateId root="2.16.840.1.113883.10.20.24.2.3"/>
          <!-- This is the LOINC code for "Measure document". This stays the same for all measure section required by QRDA standard -->
          <code code="55186-1" codeSystem="2.16.840.1.113883.6.1"/>
          <title>Measure Section</title>
          <text>
            <table border="1" width="100%">
              <thead>
                <tr>
                  <th>eMeasure Title</th>
                  <th>Version neutral identifier</th>
                  <th>eMeasure Version Number</th>
                  <th>NQF eMeasure Number</th>
                  <th>Version specific identifier</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>Initiation and Engagement of Alcohol and Other Drug Dependence Treatment</td>
                  <td>C3657D72-21B4-4675-820A-86C7FE293BF5</td>
                  <td>1</td>
                  <td>50d3a284da5fe6e140000039</td>
                  <td>8A4D92B2-3946-CDAE-0139-7944ACB700BD</td>
                  <td>
                  </td>
                </tr>
              </tbody>
            </table>
          </text>
          <!-- 1..* Organizers, each containing a reference to an eMeasure -->
          <entry>
            <organizer classCode="CLUSTER" moodCode="EVN">
              <!-- This is the templateId for Measure Reference -->
              <templateId root="2.16.840.1.113883.10.20.24.3.98"/>
              <!-- This is the templateId for eMeasure Reference QDM -->
              <templateId root="2.16.840.1.113883.10.20.24.3.97"/>
              <id extension="F7EE2890-2A44-0132-59D9-7CD1C3F6890B" />
              <statusCode code="completed"/>
              <!-- Containing isBranch external references -->
              <reference typeCode="REFR">
                <externalDocument classCode="DOC" moodCode="EVN">
                  <!-- SHALL: This is the version specific identifier for eMeasure: QualityMeasureDocument/id it is a GUID-->
                  <id root="2.16.840.1.113883.4.738" extension="8A4D92B2-3946-CDAE-0139-7944ACB700BD"/>
                  <!-- SHOULD This is the title of the eMeasure -->
                  <text>Initiation and Engagement of Alcohol and Other Drug Dependence Treatment</text>
                  <!-- SHOULD: setId is the eMeasure version neutral id  -->
                  <setId root="C3657D72-21B4-4675-820A-86C7FE293BF5"/>
                  <!-- This is the sequential eMeasure Version number -->
                  <versionNumber value="1"/>
                </externalDocument>
              </reference>
            </organizer>
          </entry>
        </section>
      </component>
      <component>
        <section>
          <!-- This is the templateId for Reporting Parameters section -->
          <templateId root="2.16.840.1.113883.10.20.17.2.1"/>
          <code code="55187-9" codeSystem="2.16.840.1.113883.6.1"/>
          <title>Reporting Parameters</title>
          <text>
            <list>
              <item>Reporting period: January 1st, 2010 00:00 - January 1st, 2011 00:00</item>
            </list>
          </text>
          <entry typeCode="DRIV">
            <act classCode="ACT" moodCode="EVN">
              <!-- This is the templateId for Reporting Parameteres Act -->
              <templateId root="2.16.840.1.113883.10.20.17.3.8"/>
              <id extension="E416F230-2A44-0132-59D9-7CD1C3F6890B" />
              <code code="252116004" codeSystem="2.16.840.1.113883.6.96" displayName="Observation Parameters"/>
              <effectiveTime>
                <low value="20100101000000"/>
                <high value="20110101000000"/>
              </effectiveTime>
            </act>
          </entry>
        </section>
      </component>
      <component>
        <section>
          <!-- This is the templateId for Patient Data section -->
          <templateId root="2.16.840.1.113883.10.20.17.2.4"/>
          <!-- This is the templateId for Patient Data QDM section -->
          <templateId root="2.16.840.1.113883.10.20.24.2.1"/>
          <code code="55188-7" codeSystem="2.16.840.1.113883.6.1"/>
          <title>Patient Data</title>
          <text>
          </text>
          <entry>
            <!-- Patient Characteristic Payer -->
            <observation classCode="OBS" moodCode="EVN">
              <templateId root="2.16.840.1.113883.10.20.24.3.55"/>
              <id root="1.3.6.1.4.1.115" extension="4"/>
              <code code="48768-6" codeSystemName="LOINC" codeSystem="2.16.840.1.113883.6.1" displayName="Payment source"/>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value="20050326153000"/>
              </effectiveTime>
              <value code="349" codeSystem="2.16.840.1.113883.3.221.5" xsi:type="CD" sdtc:valueSet="2.16.840.1.114222.4.11.3591">
                <originalText>
                </originalText>
              </value>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="RQO">
              <!-- Consolidated Plan of Care Activity Observation                   templateId (Implied Template) -->
              <templateId root="2.16.840.1.113883.10.20.22.4.44"/>
              <!-- Diagnostic Study, Order template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.17"/>
              <id root="50f84dbb7042f9366f00014c"/>
              <code code="69399-4" codeSystem="2.16.840.1.113883.6.1" sdtc:valueSet="2.16.840.1.113883.3.117.1.7.1.276">
                <originalText>Diagnostic Study, Order: VTE Diagnostic Test (Code List: 2.16.840.1.113883.3.117.1.7.1.276)</originalText>
              </code>
              <text>Diagnostic Study, Order: VTE Diagnostic Test (Code List: 2.16.840.1.113883.3.117.1.7.1.276)</text>
              <statusCode code="new"/>
              <!-- Attribute: datetime -->
              <author>
                <time value='19891215072420'/>
                <assignedAuthor>
                  <id nullFlavor="NA" root="2107e410-4308-0130-0ade-680688cbd736"/>
                </assignedAuthor>
              </author>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!--  Problem observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.4"/>
              <!--  Diagnosis active template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.11"/>
              <id root="1.3.6.1.4.1.115" extension="54c1142869702d2cd2520100"/>
              <code code="282291009" displayName="diagnosis" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED-CT"/>
              <text>Diagnosis, Active: Atrial Fibrillation/Flutter</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <!-- Attribute: admission datetime -->
                <low value='20120326153000'/>
                <!-- Attribute: discharge datetime -->
                <high nullFlavor='UNK'/>
              </effectiveTime>
              <value code="195080001" codeSystem="2.16.840.1.113883.6.96" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.117.1.7.1.202">
                <originalText>Diagnosis, Active: Atrial Fibrillation/Flutter</originalText>
                <translation code="427.31" codeSystem="2.16.840.1.113883.6.103"/>
                <translation code="I48.0" codeSystem="2.16.840.1.113883.6.90"/>
              </value>
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Problem Status (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.6"/>
                  <!-- Problem Status, Active template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.94"/>
                  <id root="7db522a0-19c6-0133-26ad-1274d76bdb3b"/>
                  <code code="33999-4"          codeSystem="2.16.840.1.113883.6.1"          codeSystemName="LOINC"          displayName="status"/>
                  <statusCode code="completed"/>
                  <value xsi:type="CD" code="55561003"           displayName="active"           codeSystem="2.16.840.1.113883.6.96"           codeSystemName="SNOMED CT"/>
                </observation>
              </entryRelationship>
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.22.4.8"/>
                  <code code="SEV"                        codeSystem="2.16.840.1.113883.5.4"/>
                  <statusCode code="completed"/>
                  <value xsi:type="CD" code="55561003" displayName="active" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT"/>
                </observation>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!--  Problem observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.4"/>
              <!--  Diagnosis active template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.11"/>
              <id root="1.3.6.1.4.1.115" extension="54c1142969702d2cd2cd0200"/>
              <code code="282291009" displayName="diagnosis" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED-CT"/>
              <text>Diagnosis, Active: Pregnancy Dx</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <!-- Attribute: admission datetime -->
                <low value='20130301150000'/>
                <!-- Attribute: discharge datetime -->
                <high value='20131020160000'/>
              </effectiveTime>
              <value code="237244005" codeSystem="2.16.840.1.113883.6.96" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.526.3.378">
                <originalText>Diagnosis, Active: Pregnancy Dx</originalText>
                <translation code="V22.0" codeSystem="2.16.840.1.113883.6.103"/>
                <translation code="Z34.00" codeSystem="2.16.840.1.113883.6.90"/>
              </value>
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Problem Status (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.6"/>
                  <!-- Problem Status, Active template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.94"/>
                  <id root="82ba54c0-19c6-0133-26ad-1274d76bdb3b"/>
                  <code code="33999-4"          codeSystem="2.16.840.1.113883.6.1"          codeSystemName="LOINC"          displayName="status"/>
                  <statusCode code="completed"/>
                  <value xsi:type="CD" code="55561003"           displayName="active"           codeSystem="2.16.840.1.113883.6.96"           codeSystemName="SNOMED CT"/>
                </observation>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!--  Problem observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.4"/>
              <!--  Diagnosis active template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.11"/>
              <id root="1.3.6.1.4.1.115" extension="54c1142869702d2cd2760100"/>
              <code code="282291009" displayName="diagnosis" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED-CT"/>
              <text>Diagnosis, Active: Diabetes</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <!-- Attribute: admission datetime -->
                <low value='20130226150000'/>
                <!-- Attribute: discharge datetime -->
                <high nullFlavor='UNK'/>
              </effectiveTime>
              <value code="46635009" codeSystem="2.16.840.1.113883.6.96" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.464.1003.103.12.1001">
                <originalText>Diagnosis, Active: Diabetes</originalText>
                <translation code="250.01" codeSystem="2.16.840.1.113883.6.103"/>
                <translation code="E10.9" codeSystem="2.16.840.1.113883.6.90"/>
              </value>
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Problem Status (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.6"/>
                  <!-- Problem Status, Active template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.94"/>
                  <id root="77791110-19c6-0133-26ad-1274d76bdb3b"/>
                  <code code="33999-4"          codeSystem="2.16.840.1.113883.6.1"          codeSystemName="LOINC"          displayName="status"/>
                  <statusCode code="completed"/>
                  <value xsi:type="CD" code="55561003"           displayName="active"           codeSystem="2.16.840.1.113883.6.96"           codeSystemName="SNOMED CT"/>
                </observation>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <encounter classCode="ENC" moodCode="EVN">
              <!--  Encounter activities template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.49"/>
              <!-- Encounter performed template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.23"/>
              <id root="1.3.6.1.4.1.115" extension="50d3a288da5fe6e14000016c"/>
              <code code="99201" codeSystem="2.16.840.1.113883.6.12" sdtc:valueSet="2.16.840.1.113883.3.464.1003.101.12.1001">
                <originalText>Encounter, Performed: Psych Visit - Individual Outpatient Psychotherapy (Code List: 2.16.840.1.113883.3.526.3.1014)</originalText>
              </code>
              <text>Encounter, Performed: Psych Visit - Individual Outpatient Psychotherapy (Code List: 2.16.840.1.113883.3.526.3.1014)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <!-- We try to look for the admit/discharge times on the encounter if they are           there. If not, we fall back to the typical start/end date. -->
                <low value='20101101120000'/>
                <high value='20101101130000'/>
              </effectiveTime>
            </encounter>
          </entry>
          <entry>
            <encounter classCode="ENC" moodCode="RQO">
              <!--  Plan of Care Activity Encounter template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.40"/>
              <!-- Encounter order template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.22"/>
              <id root="50f84c1b7042f9877500025e"/>
              <code nullFlavor="UNK" >
                <originalText>Encounter, Order: BH Outpatient psychotherapy (Code List: 2.16.840.1.113883.3.1257.1.973)</originalText>
                <translation code="76168009" codeSystem="2.16.840.1.113883.6.96"/>
                <translation code="90815" codeSystem="2.16.840.1.113883.6.12"/>
                <translation code="94.49" codeSystem="2.16.840.1.113883.6.103"/>
                <translation code="GZHZZZZ" codeSystem="2.16.840.1.113883.6.4"/>
              </code>
              <text>Encounter, Order: BH Outpatient psychotherapy (Code List: 2.16.840.1.113883.3.1257.1.973)</text>
              <statusCode code="new"/>
              <!-- Attribute: datetime -->
              <author>
                <time value='20051226144034'/>
                <assignedAuthor>
                  <id root="292627f0-4307-0130-0add-680688cbd736"/>
                </assignedAuthor>
              </author>
            </encounter>
          </entry>
          <entry>
            <encounter classCode="ENC" moodCode="EVN">
              <!--  Encounter activities template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.49"/>
              <!-- Encounter performed template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.23"/>
              <id root="1.3.6.1.4.1.115" extension="50d3a288da5fe6e14000016a"/>
              <code code="99201" codeSystem="2.16.840.1.113883.6.12" sdtc:valueSet="2.16.840.1.113883.3.464.1003.101.12.1001">
                <originalText>Encounter, Performed: Office Visit (Code List: 2.16.840.1.113883.3.464.1003.101.12.1001)</originalText>
              </code>
              <text>Encounter, Performed: Office Visit (Code List: 2.16.840.1.113883.3.464.1003.101.12.1001)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <!-- We try to look for the admit/discharge times on the encounter if they are           there. If not, we fall back to the typical start/end date. -->
                <low value='20091101000000'/>
                <high value='20091101000000'/>
              </effectiveTime>
            </encounter>
          </entry>
          <entry>
            <encounter classCode="ENC" moodCode="EVN">
              <!--  Encounter activities template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.49"/>
              <!-- Encounter performed template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.23"/>
              <id root="1.3.6.1.4.1.115" extension="50d3a288da5fe6e14000016b"/>
              <code code="99201" codeSystem="2.16.840.1.113883.6.12" sdtc:valueSet="2.16.840.1.113883.3.464.1003.101.12.1001">
                <originalText>Encounter, Performed: Office Visit (Code List: 2.16.840.1.113883.3.464.1003.101.12.1001)</originalText>
              </code>
              <text>Encounter, Performed: Office Visit (Code List: 2.16.840.1.113883.3.464.1003.101.12.1001)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <!-- We try to look for the admit/discharge times on the encounter if they are           there. If not, we fall back to the typical start/end date. -->
                <low value='20101001120000'/>
                <high value='20101001130000'/>
              </effectiveTime>
            </encounter>
          </entry>
          <entry>
            <substanceAdministration classCode="SBADM" moodCode="EVN">
              <!-- Conforms to C-CDA R2 Medication Activity (V2) -->
              <templateId root="2.16.840.1.113883.10.20.22.4.16" extension="2014-06-09" />
              <!-- Medication Active (V2) template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.41" extension="2014-12-01" />
              <id root="c0ea7bf3-50e7-4e7a-83a3-e5a9ccbb8541" />
              <statusCode code="active" />
              <effectiveTime xsi:type="IVL_TS">
                <!-- QDM Attribute: Start datetime -->
                <low value="20040816121859" />
                <!-- QDM Attribute: Stop datetime -->
                <high value="20040816160606" />
              </effectiveTime>
              <!-- QDM Attribute: Frequency -->
              <effectiveTime xsi:type="PIVL_TS" institutionSpecified="true" operator="A">
                <period value="6" unit="h" />
              </effectiveTime>
              <maxDoseQuantity nullFlavor="UNK">
                <numerator value="42" unit="oz" />
                <denominator value="100" unit="oz" />
              </maxDoseQuantity>
              <administrationUnitCode code="C42944" displayName="INHALANT" codeSystem="2.16.840.1.113883.3.26.1.1" codeSystemName="NCI Thesaurus"/>
              <!-- QDM Attribute: Route -->
              <routeCode code="C38288"        codeSystem="2.16.840.1.113883.3.26.1.1"        codeSystemName="NCI Thesaurus"        displayName="ORAL" />
              <!-- QDM Attribute: Dose -->
              <doseQuantity value="1"/>
              <consumable>
                <manufacturedProduct classCode="MANU">
                  <!-- Conforms to C-CDA R2 Medication Information (V2) -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.23" extension="2014-06-09" />
                  <id root="37bfe02a-3e97-4bd6-9197-bbd0ed0de79e" />
                  <manufacturedMaterial>
                    <code code="105152" codeSystem="2.16.840.1.113883.6.88"           codeSystemName="RxNorm"           displayName="Amoxicillin 60 MG/ML Oral Suspension"           sdtc:valueSet="{$QDMElementValueSetOID}" />
                  </manufacturedMaterial>
                </manufacturedProduct>
              </consumable>
              <entryRelationship typeCode="REFR">
                <supply classCode="SPLY" moodCode="INT">
                  <templateId root="2.16.840.1.113883.10.20.22.4.17"/>
                  <id root="12345"/>
                  <statusCode code="completed"/>
                  <effectiveTime xsi:type="IVL_TS">
                    <low value="20040816160606"/>
                    <high nullFlavor="UNK"/>
                  </effectiveTime>
                  <repeatNumber value="1"/>
                  <quantity value="75"/>
                  <author>
                    <time nullFlavor="UNK"/>
                    <assignedAuthor>
                      <id/>
                      <addr nullFlavor="UNK"/>
                      <telecom nullFlavor="UNK"/>
                      <assignedPerson>
                        <name>
                          <prefix>Dr.</prefix>
                          <given>Robert</given>
                          <family>Michaels</family>
                        </name>
                      </assignedPerson>
                    </assignedAuthor>
                  </author>
                  <entryRelationship typeCode="SUBJ" inversionInd="true">
                    <act classCode="ACT" moodCode="INT">
                      <templateId root="2.16.840.1.113883.10.20.22.4.20"/>
                      <!-- ** Instructions Template ** -->
                      <code code="409073007" codeSystem="2.16.840.1.113883.6.96" displayName="instruction"/>
                      <text>label in spanish</text>
                      <statusCode code="completed"/>
                    </act>
                  </entryRelationship>
                </supply>
              </entryRelationship>
            </substanceAdministration>
          </entry>
          <entry>
            <!--Medication dispensed -->
            <supply classCode="SPLY" moodCode="EVN">
              <!--  Medication Dispensed template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.45"/>
              <!-- Medication Dispense template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.18"/>
              <id root="50f84c1b7042f9877500023e"/>
              <text>Medication, Dispensed: ADHD Medications (Code List: 2.16.840.1.113883.3.464.1003.196.12.1171)</text>
              <statusCode code="completed"/>
              <effectiveTime xsi:type="IVL_TS">
                <low value='19960119172123'/>
                <high value='19960119221325'/>
              </effectiveTime>
              <product>
                <manufacturedProduct classCode="MANU">
                  <!-- Medication Information (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.23"/>
                  <id root="2914e300-4307-0130-0add-680688cbd736"/>
                  <manufacturedMaterial>
                    <code nullFlavor="UNK" >
                      <originalText>Medication, Dispensed: ADHD Medications (Code List: 2.16.840.1.113883.3.464.1003.196.12.1171)</originalText>
                      <translation code="977869" codeSystem="2.16.840.1.113883.6.88"/>
                    </code>
                  </manufacturedMaterial>
                </manufacturedProduct>
              </product>
            </supply>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!-- consolidation CDA Allergy observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.7"/>
              <!--  Medication Allergy -->
              <templateId root="2.16.840.1.113883.10.20.24.3.43"/>
              <id root="50f84db97042f9366f00000e"/>
              <code code="ASSERTION" displayName="Assertion" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='19790809140056'/>
              </effectiveTime>
              <value code="62014003" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Adverse drug effect" xsi:type="CD"/>

              <participant typeCode="CSM">
                <participantRole classCode="MANU">
                  <playingEntity classCode="MMAT">
                    <code code="996994" codeSystem="2.16.840.1.113883.6.88" sdtc:valueSet="2.16.840.1.113883.3.666.5.626"><originalText>Medication, Adverse Effects: Hospital Measures-Aspirin (Code List: 2.16.840.1.113883.3.666.5.626)</originalText></code>
                    <name>Medication, Adverse Effects: Hospital Measures-Aspirin (Code List: 2.16.840.1.113883.3.666.5.626)</name>
                  </playingEntity>
                </participantRole>
              </participant>
              <!-- QDM Attribute: Reaction -->
              <entryRelationship typeCode="MFST" inversionInd="true">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Conforms to C-CDA R2 Reaction Observation (V2) -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.9" extension="2014-06-09" />
                  <!-- Reaction (QRDA I V2) -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.85" extension="2014-12-01" />
                  <value xsi:type="CD" code="422587007" codeSystem="2.16.840.1.113883.6.96" displayName="Nausea" />
                </observation>
              </entryRelationship>
              <entryRelationship typeCode="SUBJ">
                <observation classCode="OBS" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.22.4.8" extension="2014-06-09" />
                  <code code="SEV" displayName="Severity Observation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode" />
                  <text>
                      <reference value="#allergyseverity1" />
                  </text>
                  <statusCode code="completed" />
                  <value xsi:type="CD" code="371924009" displayName="Moderate to severe" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" />
              </observation>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <templateId root="2.16.840.1.113883.10.20.24.3.62" extension="2014-12-01"/>
              <templateId root="2.16.840.1.113883.10.20.24.3.104" extension="2014-12-01"/>
              <id root="5102936b944dfe3db4000016"/>
              <code code="ASSERTION"
                codeSystem="2.16.840.1.113883.5.4"
                codeSystemName="ActCode"
                displayName="Assertion"/>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value="20040912123835"/>
              </effectiveTime>
              <value xsi:type="CD"
              code="102460003"
              codeSystem="2.16.840.1.113883.6.96"
              codeSystemName="SNOMED CT"
              displayName="Decreased tolerance"/>
              <entryRelationship typeCode="CAUS" inversionInd="true">
                <procedure classCode="PROC" moodCode="EVN">
                  <!--  Procedure performed template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.64" extension="2014-12-01"/>
                  <!-- Procedure Activity Procedure-->
                  <templateId root="2.16.840.1.113883.10.20.22.4.14" extension="2014-12-01"/>
                  <id root="5102936b944dfe3db4000016"/>
                  <code code="86198006" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.526.3.402"><originalText>Procedure, Intolerance: Influenza Vaccination (Code List: 2.16.840.1.113883.3.526.3.402)</originalText><translation code="90668" codeSystem="2.16.840.1.113883.6.12"/>
                  </code>
                  <text>Procedure, Intolerance: Influenza Vaccination (Code List: 2.16.840.1.113883.3.526.3.402)</text>
                  <statusCode code="completed"/>
                  <effectiveTime>
                    <low value='20040912123835'/>
                    <high value='20040913023209'/>
                  </effectiveTime>
                </procedure>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!--  Problem observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.4"/>
              <!-- Diagnosis, inactive template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.13"/>
              <id root="50f84c1d7042f98775000352"/>
              <code code="282291009" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="diagnosis"/>
              <statusCode code="completed"/>
              <effectiveTime>
                <!-- Attribute: Start Datetime -->
                <low value='20040816121859'/>
                <!-- Attribute: Stop Datetime -->
                <high value='20040816200929'/>
              </effectiveTime>
              <value code="76795007" codeSystem="2.16.840.1.113883.6.96" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.67.1.101.1.269">
                <originalText>Diagnosis, Inactive: Hepatitis B (Code List: 2.16.840.1.113883.3.67.1.101.1.269)</originalText>
                <translation code="V02.61" codeSystem="2.16.840.1.113883.6.103"/>
                <translation code="Z22.51" codeSystem="2.16.840.1.113883.6.90"/>
              </value>
              <!-- Status -->
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Problem Status (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.6"/>
                  <!-- Problem Status, Inactive template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.95"/>
                  <id root="2a0bd200-4307-0130-0add-680688cbd736"/>
                  <code code="33999-4" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" displayName="status"/>
                  <statusCode code="completed"/>
                  <value code="73425007" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="inactive" xsi:type="CD"/>
                </observation>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="RQO">
              <!-- Consolidation Plan of Care Activity Observation -->
              <templateId root="2.16.840.1.113883.10.20.22.4.44"/>
              <!-- Lab Test Order -->
              <templateId root="2.16.840.1.113883.10.20.24.3.37"/>
              <id root="50f84c1d7042f9877500039e"/>
              <code nullFlavor="UNK" sdtc:valueSet="2.16.840.1.113883.3.600.1482">
                <originalText>Laboratory Test, Order: Laboratory Tests for Hypertension (Code List: 2.16.840.1.113883.3.600.1482)</originalText>
                <translation code="8879006" codeSystem="2.16.840.1.113883.6.96"/>
                <translation code="80069" codeSystem="2.16.840.1.113883.6.12"/>
              </code>
              <text>Laboratory Test, Order: Laboratory Tests for Hypertension (Code List: 2.16.840.1.113883.3.600.1482)</text>
              <statusCode code="new"/>
              <!-- Attribute: datetime -->
              <author>
                <time value='19910519162436'/>
                <assignedAuthor>
                  <id nullFlavor="NA" root="2a5d88e0-4307-0130-0add-680688cbd736"/>
                </assignedAuthor>
              </author>
            </observation>
          </entry>
          <entry>
            <!--Laboratory test, result -->
            <observation classCode="OBS" moodCode="EVN">
              <!--  Result Observation (consolidation) template  -->
              <templateId root="2.16.840.1.113883.10.20.22.4.2"/>
              <!-- Laboratory Test, Result template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.40"/>
              <id root="1.3.6.1.4.1.115" extension="50d3a288da5fe6e1400002a9"/>
              <code code="11268-0" codeSystem="2.16.840.1.113883.6.1" sdtc:valueSet="2.16.840.1.113883.3.464.1003.198.12.1012">
                <originalText>Laboratory Test, Result: Group A Streptococcus Test (Code List: 2.16.840.1.113883.3.464.1003.198.12.1012)</originalText>
              </code>
              <text>Laboratory Test, Result: Group A Streptococcus Test (Code List: 2.16.840.1.113883.3.464.1003.198.12.1012)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value="19910519162436"/>
              </effectiveTime>
              <value xsi:type="ST" >positive</value>
            </observation>
          </entry>
          <entry>
            <encounter classCode="ENC" moodCode="EVN">
              <!-- Transfer From -->
              <templateId root="2.16.840.1.113883.10.20.24.3.81" extension="2014-12-01" />
              <id root="49d75f61-0dec-4972-9a51-e2490b18c772" />
              <code code="77305-1" codeSystem="2.16.840.1.113883.6.1" displayName="Transferred from" codeSystemName="LOINC" />
              <statusCode code="completed" />
              <participant typeCode="ORG">
                <time>
                  <!-- QDM Attribute: Start Datetime -->
                  <low value="20141104103000" />
                  <!-- SHALL have low -->
                  <!-- QDM Attribute: Start Datetime -->
                  <high value="20141104103000" />
                  <!-- MAY have high -->
                </time>
                <participantRole classCode="LOCE">
                  <!-- Transfer from location -->
                  <code code="309911002" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED-CT" displayName="Respiratory intensive care unit (environment)" sdtc:valueSet="{$QDMElementValueSetOID}" />
                </participantRole>
              </participant>
            </encounter>
          </entry>
          <entry>
            <encounter classCode="ENC" moodCode="EVN">
              <!-- Transfer To -->
              <templateId root="2.16.840.1.113883.10.20.24.3.82" extension="2014-12-01" />
              <id root="49d75f61-0dec-4972-9a51-e2490b18c772" />
              <code code="77306-9" codeSystem="2.16.840.1.113883.6.1" displayName="Discharge disposition" codeSystemName="LOINC" />
              <statusCode code="completed" />
              <participant typeCode="DST">
                <time>
                  <!-- QDM Attribute: Start Datetime -->
                  <low value="20141104103000" />
                  <!-- SHALL have low -->
                  <!-- QDM Attribute: Start Datetime -->
                  <high value="20141104103000" />
                  <!-- MAY have high -->
                </time>
                <participantRole classCode="LOCE">
                  <!-- Transfer to location -->
                  <code code="309911002" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED-CT" displayName="Respiratory intensive care unit (environment)" sdtc:valueSet="{$QDMElementValueSetOID}" />
                </participantRole>
              </participant>
            </encounter>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <templateId root="2.16.840.1.113883.10.20.24.3.101"/>
              <id root="50f6c6da7042f9cdd0000233"/>
              <code code="57036006" codeSystem="2.16.840.1.113883.6.96"
                    codeSystemName="SNOMED-CT" displayName="length of gestation"/>
              <statusCode code="completed"/>
              <effectiveTime value="19920812064214"/>
              <!-- Note: valueSet of PQ "d" and "wk"-->
              <value xsi:type="PQ" value="36" unit="wk"/>
            </observation>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="EVN">
              <!-- Communication from patient to provider -->
              <templateId root="2.16.840.1.113883.10.20.24.3.2"/>
              <id root="50f84c187042f987750000e5"/>
              <code nullFlavor="UNK" ><originalText>Communication: From Patient to Provider: Influenza Vaccination Declined (Code List: 2.16.840.1.113883.3.526.3.1255)</originalText><translation code="315640000" codeSystem="2.16.840.1.113883.6.96"/>
              </code>
              <text>
                 Communication: From Patient to Provider: Influenza Vaccination Declined (Code List: 2.16.840.1.113883.3.526.3.1255)
              </text>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='20101027165345'/>
              </effectiveTime>

              <participant typeCode="AUT">
                <participantRole classCode="PAT">
                  <code code="116154003" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Patient"/>
                </participantRole>
              </participant>

              <participant typeCode="IRCP">
                <participantRole classCode="ASSIGNED">
                  <code code="158965000" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Medical Practitioner"/>
                </participantRole>
              </participant>

              <entryRelationship typeCode="RSON">
                <observation classCode="OBS" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.24.3.88"/>
                  <id extension="5B2220FB0B57C9CD39A9AE22C835ACC6"/>
                  <code code="410666004" codeSystem="2.16.840.1.113883.6.96" displayName="reason" codeSystemName="SNOMED CT"/>
                  <statusCode code="completed"/>
                  <effectiveTime value="20101027165345"/>
                  <value xsi:type="CD" code="105480006" codeSystem="2.16.840.1.113883.6.96"/>
                </observation>
              </entryRelationship>

              <sdtc:inFulfillmentOf1 typeCode="FLFS">
                <sdtc:templateId root="2.16.840.1.113883.10.20.24.3.126" extension="2014-12-01" />
                <sdtc:actReference classCode="ACT" moodCode="EVN">
                  <sdtc:id root="1.3.6.1.4.1.115" extension="56c237ee02d40565bb00030e"/>
                </sdtc:actReference>
              </sdtc:inFulfillmentOf1>

            </act>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="EVN" >
              <!-- Consolidation CDA: Procedure Activity Act template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.12"/>
              <templateId root="2.16.840.1.113883.10.20.24.3.32"/>
              <id root="510831719eae47faed00019f"/>
              <code code="171207006" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.600.426"><originalText>Intervention, Performed: Adolescent Depression Screening (Code List: 2.16.840.1.113883.3.600.426)</originalText><translation code="S3005" codeSystem="2.16.840.1.113883.6.14"/>
                <translation code="V79.0" codeSystem="2.16.840.1.113883.6.103"/>
                <translation code="Z13.89" codeSystem="2.16.840.1.113883.6.90"/>
              </code>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='20100205120000'/>
                <high value='20100205120000'/>
              </effectiveTime>
            </act>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <templateId root="2.16.840.1.113883.10.20.24.3.103"/>
              <id root="50f6c6067042f91c7c000272"/>
              <code code="ASSERTION" codeSystem="2.16.840.1.113883.5.4"/>
              <statusCode code="completed"/>
              <effectiveTime value="20120619085355"/>
              <value code="423237006" codeSystem="2.16.840.1.113883.6.96" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.526.3.1279">
                <originalText>Patient Characteristic: ECOG Performance Status-Poor (Code List: 2.16.840.1.113883.3.526.3.1279)</originalText>
              </value>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!--  Problem observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.4"/>
              <!-- Symptom, active template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.76"/>
              <id root="50f84dbb7042f9366f0001ac"/>
              <code code="418799008" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Symptom"/>
              <text>Symptom, Active: Neurologic Symptoms of Stroke (Code List: 2.16.840.1.113883.3.117.1.7.1.399)</text>
              <statusCode code="completed"/>
              <value code="95815000" codeSystem="2.16.840.1.113883.6.96" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.117.1.7.1.399"><originalText>Symptom, Active: Neurologic Symptoms of Stroke (Code List: 2.16.840.1.113883.3.117.1.7.1.399)</originalText></value>
              <effectiveTime>
                <low value='19930215222215'/>
                <high value='19930216125308'/>
              </effectiveTime>
              <!-- Status -->
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Problem Status (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.6"/>
                  <!-- Problem Status, Active template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.94"/>
                  <id root="213a6870-4308-0130-0ade-680688cbd736"/>
                  <code code="33999-4" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" displayName="status"/>
                  <statusCode code="completed"/>
                  <value code="55561003" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="active" xsi:type="CD"/>
                </observation>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!--  Problem observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.4"/>
              <!-- Diagnosis, resolved template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.14"/>
              <id root="50f84c187042f98775000089"/>
              <code code="282291009" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="diagnosis"/>

              <statusCode code="completed"/>
              <effectiveTime>
                <!-- Attribute: Start Datetime -->
                <low value='20030403010254'/>
                <!-- Attribute: Stop Datetime -->
                <high value='20030403121452'/>
              </effectiveTime>
              <value code="94643001" codeSystem="2.16.840.1.113883.6.96" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.464.1003.108.12.1001"><originalText>Diagnosis, Resolved: Malignant Neoplasm of Colon (Code List: 2.16.840.1.113883.3.464.1003.108.12.1001)</originalText><translation code="197.5" codeSystem="2.16.840.1.113883.6.103"/>
              <translation code="C21.8" codeSystem="2.16.840.1.113883.6.90"/>
              </value>
              <!-- Status -->
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Problem Status (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.6"/>
                  <!-- Problem Status, Resolved template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.96"/>
                  <id root="272e2b20-4307-0130-0add-680688cbd736"/>
                  <code code="33999-4" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" displayName="status"/>
                  <statusCode code="completed"/>
                  <value code="413322009" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="resolved" xsi:type="CD"/>
                </observation>
              </entryRelationship>
            </observation>
          </entry>
          <!-- procedure with negationInd and no reason -->
          <entry>
            <observation classCode="OBS" moodCode="EVN" >
              <!-- Lab test performed -->
              <templateId root="2.16.840.1.113883.10.20.24.3.38"/>
              <id root="50f84c1d7042f98775000353"/>
              <code code="7905-3" codeSystem="2.16.840.1.113883.6.1" sdtc:valueSet="2.16.840.1.113883.3.67.1.101.1.279"><originalText>Laboratory Test, Performed: HBsAg (Code List: 2.16.840.1.113883.3.67.1.101.1.279)</originalText></code>
              <text>Laboratory Test, Performed: HBsAg (Code List: 2.16.840.1.113883.3.67.1.101.1.279)</text>
              <statusCode code="completed"/>

              <effectiveTime>
                <low value='20020129180704'/>
                <high value='20020130074815'/>
              </effectiveTime>
              <interpretationCode code="N" codeSystem="2.16.840.1.113883.5.83"/>
              <referenceRange>
                <observationRange>
                  <text>M 13-18 g/dl; F 12-16 g/dl</text>
                </observationRange>
              </referenceRange>

              <entryRelationship typeCode="RSON">
                <observation classCode="OBS" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.24.3.88"/>
                  <id extension="5B2220FB0B57C9CD39A9AE22C835ACC6"/>
                  <code code="410666004" codeSystem="2.16.840.1.113883.6.96" displayName="reason" codeSystemName="SNOMED CT"/>
                  <statusCode code="completed"/>
                  <effectiveTime value="20101027165345"/>
                  <value xsi:type="CD" code="105480006" codeSystem="2.16.840.1.113883.6.96"/>
                </observation>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <procedure classCode="PROC" moodCode="EVN" >
              <!-- Procedure Activity Procedure -->
              <templateId root="2.16.840.1.113883.10.20.22.4.14"/>
              <!-- Device Applied -->
              <templateId root="2.16.840.1.113883.10.20.24.3.7"/>
              <id root="510969b3944dfe9bd7000056"/>
              <code code="360030002" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="application of device"/>
              <text>Device, Applied</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='19850331043808'/>
              </effectiveTime>
              <targetSiteCode code="thigh"
                    codeSystem="2.16.840.1.113883.6.96"/>
              <participant typeCode="DEV">
                <participantRole classCode="MANU">
                  <playingDevice classCode="DEV">
                    <code nullFlavor="UNK" sdtc:valueSet="2.16.840.1.113883.3.666.5.1086"><originalText>Device, Applied: Hospital measures-Pacemaker or implantable defibrillator device (Code List: 2.16.840.1.113883.3.666.5.1086)</originalText><translation code="37.98" codeSystem="2.16.840.1.113883.6.103"/>
                    </code>
                  </playingDevice>
                </participantRole>
              </participant>
            </procedure>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="RQO" negationInd="true">
                <id root="1.3.6.1.4.1.115" extension="123456"/>
                <code code="SPLY" codeSystem="2.16.840.1.113883.5.6"/>
                <text>Medical Device Not Ordered</text>
                <statusCode code="active"/>
                <effectiveTime xsi:type="IVL_TS">
                  <low value="20100101000000"/>
                  <high value="20110101000000"/>
                </effectiveTime>
                <author>
                  <templateId root="2.16.840.1.113883.10.20.22.4.119"/>
                  <time value="20100101000000"/>
                  <assignedAuthor>
                      <id nullFlavor="NA" root="2o141ol3j4121m30238ht5q3u4n15"/>
                  </assignedAuthor>
                </author>
                <participant typeCode="DEV">
                  <participantRole classCode="MANU">
                    <playingDevice classCode="DEV">
                      <code nullFlavor="UNK" sdtc:valueSet="2.16.840.1.113883.3.666.5.1086"><originalText>Device Not Ordered: Hospital measures-Pacemaker or implantable defibrillator device (Code List: 2.16.840.1.113883.3.666.5.1086)</originalText><translation code="48.20" codeSystem="2.16.840.1.113883.6.103"/>
                      </code>
                    </playingDevice>
                  </participantRole>
                </participant>
              </act>
            </entry>
          <entry>
            <procedure classCode="PROC" moodCode="EVN" negationInd="true">
              <!--  Procedure performed template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.64"/>
              <!-- Procedure Activity Procedure-->
              <templateId root="2.16.840.1.113883.10.20.22.4.14"/>
              <id root="51083f0e944dfe9bd7000004"/>
              <code code="236211007" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.526.3.399"><originalText>Procedure, Performed: Salvage Therapy (Code List: 2.16.840.1.113883.3.526.3.399)</originalText><translation code="55876" codeSystem="2.16.840.1.113883.6.12"/>
              </code>
              <text>Procedure, Performed: Salvage Therapy (Code List: 2.16.840.1.113883.3.526.3.399)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='19860117151405'/>
                <high value='19860118051933'/>
              </effectiveTime>
              <targetSiteCode code="28273000" displayName="bile duct" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" /> <!-- used for anatimical target. found in "procedure activity procedure" example -->
              <priorityCode code="63161005" codeSystem="2.16.840.1.113883.6.96"/>
              <entryRelationship typeCode="REFR">
                <procedure classCode="PROC" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.24.3.89"/>
                  <code code="34896006" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="incision"/>
                  <effectiveTime value="19860117151505"/>
                </procedure>
              </entryRelationship>
            </procedure>
          </entry>
          <!-- procedure with negationInd and reason. contains multiple values to test different xsi types -->
          <entry>
            <procedure classCode="PROC" moodCode="EVN" negationInd="true">
              <!--  Procedure performed template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.64"/>
              <!-- Procedure Activity Procedure-->
              <templateId root="2.16.840.1.113883.10.20.22.4.14"/>
              <id root="51083f0e944dfe9bd7001234"/>
              <code code="236211007" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.526.3.399"><originalText>Procedure, Performed: Salvage Therapy (Code List: 2.16.840.1.113883.3.526.3.399)</originalText><translation code="55876" codeSystem="2.16.840.1.113883.6.12"/>
              </code>
              <text>Procedure, Performed: Salvage Therapy (Code List: 2.16.840.1.113883.3.526.3.399)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='19860117151405'/>
                <high value='19860118051933'/>
              </effectiveTime>
              <targetSiteCode code="28273000" displayName="bile duct" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" /> <!-- used for anatimical target. found in "procedure activity procedure" example. added by Jesse -->
              <priorityCode code="63161005" codeSystem="2.16.840.1.113883.6.96"/>
              <entryRelationship typeCode="REFR">
                <procedure classCode="PROC" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.24.3.89"/>
                  <code code="34896006" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="incision"/>
                  <effectiveTime value="19860117151505"/>
                </procedure>
              </entryRelationship>
              <entryRelationship typeCode="RSON">
                <act classCode="ACT" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.1.27"/> <!-- added template id to identify negation reason -->
                  <code code="308292007" codeSystem="2.16.840.1.113883.6.96" displayName="Transfer of care" codeSystemName="SNOMED CT"/>
                  <statusCode code="completed"/>
                </act>
              </entryRelationship>
              <value xsi:type="PQ" value="6" unit="m[IU]/L"/> <!-- added to test values tags with different formats -->
              <value xsi:type="BL" value="true"/>
              <value xsi:type="ST">my_string_value</value>
            </procedure>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN" negationInd="true">
              <!-- Procedure Activity Procedure (Consolidation) template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.13"/>
              <!-- Physical Exam, Performed template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.59"/>
              <id root="5101a4f7944dfe3db4000006"/>
              <code nullFlavor="UNK" sdtc:valueSet="2.16.840.1.113883.3.526.3.1033"><originalText>Physical Exam, Performed: Diastolic Blood Pressure (Code List: 2.16.840.1.113883.3.526.3.1033)</originalText><translation code="8462-4" codeSystem="2.16.840.1.113883.6.1"/>
              </code>
              <text>Physical Exam, Performed: Diastolic Blood Pressure (Code List: 2.16.840.1.113883.3.526.3.1033)</text>
              <statusCode code="completed"/>
              <effectiveTime>
              <low value='19931019040716'/>
              <high value='19931019195142'/>
              </effectiveTime>
              <value xsi:type="CD" nullFlavor="UNK"/>
            </observation>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="RQO">
              <!-- Plan of Care Activity Act template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.39"/>
              <!-- Intervention Order template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.31"/>
              <id root="510831719eae47faed000150"/>
              <code code="304549008" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.600.1.1525">
                <originalText>Intervention, Order: Above Normal Follow-up (Code List: 2.16.840.1.113883.3.600.1.1525)</originalText>
                <translation code="43644" codeSystem="2.16.840.1.113883.6.12"/>
                <translation code="G8417" codeSystem="2.16.840.1.113883.6.14"/>
                <translation code="V65.3" codeSystem="2.16.840.1.113883.6.103"/>
                <translation code="Z71.3" codeSystem="2.16.840.1.113883.6.90"/>
              </code>
              <statusCode code="completed"/>
              <author>
                <time value='20100625000000'/>
                <assignedAuthor>
                  <id root="56c10b90-4d39-0130-5b7f-0017f2071af4"/>
                </assignedAuthor>
              </author>
              <effectiveTime>
                <low value='20100625000000'/>
              </effectiveTime>
            </act>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="EVN">
              <!-- Consolidation CDA: Procedure Activity Act template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.12"/>
              <templateId root="2.16.840.1.113883.10.20.24.3.34"/>
              <id root="50f84c1c7042f987750002d1"/>
              <code code="428181000124104" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.600.793"><originalText>Intervention, Result: Positive Depression Screening (Code List: 2.16.840.1.113883.3.600.793)</originalText></code>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='20041016152724'/>
                <high value='20041016204832'/>
              </effectiveTime>
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Result Observation template (consolidation) -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.2"/>
                  <!-- Result template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.87"/>
                  <id root="2994ccb0-4307-0130-0add-680688cbd736"/>
                  <code code="428181000124104" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.600.793"><originalText>Intervention, Result: Positive Depression Screening (Code List: 2.16.840.1.113883.3.600.793)</originalText></code>
                  <statusCode code="completed"/>
                  <effectiveTime value='20041016152724'/>
                  <value xsi:type="PQ" nullFlavor="UNK"/>

                  <!-- Attribute: Status -->
                  <entryRelationship typeCode="REFR">
                    <observation classCode="OBS" moodCode="EVN">
                      <templateId root="2.16.840.1.113883.10.20.24.3.93"/>
                      <id root="2994d3b0-4307-0130-0add-680688cbd736"/>
                      <code code="33999-4"
                            codeSystem="2.16.840.1.113883.6.1"
                            codeSystemName="LOINC"
                            displayName="Status"/>
                      <value xsi:type="CD"
                            code="33999-4"
                            codeSystem="2.16.840.1.113883.6.1"
                            displayName="Final Report"/>
                    </observation>
                  </entryRelationship>
                </observation>
              </entryRelationship>
            </act>
          </entry>
          <entry>
            <procedure classCode="PROC" moodCode="EVN">
              <!-- Consolidated Procedure Activity Procedure TemplateId
                   (Implied Template) -->
              <templateId root="2.16.840.1.113883.10.20.22.4.14"/>
              <!-- QRDA Procedure, Result TemplateId -->
              <templateId root="2.16.840.1.113883.10.20.24.3.66"/>
              <id root="51095fc3944dfe9bd7000012"/>
              <code code="116783008" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.526.3.1098"><originalText>Procedure, Result: Clinical Staging Procedure (Code List: 2.16.840.1.113883.3.526.3.1098)</originalText></code>
              <text>Procedure, Result: Clinical Staging Procedure (Code List: 2.16.840.1.113883.3.526.3.1098)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='20011202034746'/>
                <high value='20011202180443'/>
              </effectiveTime>
              <entryRelationship typeCode="REFR">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Result Observation template (consolidation) -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.2"/>
                  <!-- Result template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.87"/>
                  <id root="e10b5470-4d34-0130-07fd-406c8f2fb681"/>
                  <code code="116783008" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.526.3.1098"><originalText>Procedure, Result: Clinical Staging Procedure (Code List: 2.16.840.1.113883.3.526.3.1098)</originalText></code>
                  <statusCode code="completed"/>
                  <effectiveTime value='20011202034746'/>
                  <value xsi:type="CD" nullFlavor="UNK"/>
                </observation>
              </entryRelationship>
            </procedure>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN" >
              <!-- Consolidation Assessment Scale Observation templateId -->
              <templateId root="2.16.840.1.113883.10.20.22.4.69"/>
              <!-- Risk Category Assessment -->
              <templateId root="2.16.840.1.113883.10.20.24.3.69"/>
              <id root="510963e9944dfe9bd7000047"/>
              <code code="72136-5" codeSystem="2.16.840.1.113883.6.1" sdtc:valueSet="2.16.840.1.113883.3.117.1.7.1.357"><originalText>Risk Category Assessment: VTE Risk Assessment (Code List: 2.16.840.1.113883.3.117.1.7.1.357)</originalText></code>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='19930805130208'/>
              </effectiveTime>
              <value xsi:type="INT" value="7"/>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN" >
              <!-- Consolidated Procedure Activity Observation templateId
                 (Implied Template) -->
              <templateId root="2.16.840.1.113883.10.20.22.4.13"/>
              <!-- Diagnostic Study, Performed template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.18"/>
              <id root="50f84dbb7042f9366f000143"/>
              <code nullFlavor="UNK" sdtc:valueSet="2.16.840.1.113883.3.117.1.7.1.276">
                <originalText>Diagnostic Study, Performed: VTE Diagnostic Test (Code List: 2.16.840.1.113883.3.117.1.7.1.276)</originalText>
                <translation code="69399-4" codeSystem="2.16.840.1.113883.6.1"/>
              </code>
              <text>Diagnostic Study, Performed: VTE Diagnostic Test (Code List: 2.16.840.1.113883.3.117.1.7.1.276)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <!-- Attribute: Start Datetime -->
                <low value='20081029211606'/>
                <!-- Attribute: End Datetime -->
                <high value='20081029230540'/>
              </effectiveTime>
              <value xsi:type="CD" nullFlavor="UNK"/>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN" negationInd="true">
              <!-- Consolidated Result Observation templateId (Implied Template) -->
              <templateId root="2.16.840.1.113883.10.20.22.4.2"/>
              <!-- Diagnostic Study, Result template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.20"/>
              <id root="50f84c1b7042f987750001e7"/>
              <code code="71485-7" codeSystem="2.16.840.1.113883.6.1" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.526.3.1333"><originalText>Diagnostic Study, Result: Cup to Disc Ratio (Code List: 2.16.840.1.113883.3.526.3.1333)</originalText></code>
              <text>Diagnostic Study, Result: Cup to Disc Ratio (Code List: 2.16.840.1.113883.3.526.3.1333)</text>
              <statusCode code="completed"/>

              <effectiveTime>
                <low value='19890923063243'/>
                <high value='19890923101231'/>
              </effectiveTime>
              <value xsi:type="CD" nullFlavor="UNK"/>

              <entryRelationship typeCode="RSON">
                <observation classCode="OBS" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.24.3.88"/>
                  <code code="410666004"
                        codeSystem="2.16.840.1.113883.6.96"
                        displayName="reason"
                        codeSystemName="SNOMED CT"/>
                  <statusCode code="completed"/>
                  <effectiveTime value='19890923063243'/>
                  <value xsi:type="CD"
                       code="79899007"
                       codeSystem=""/>
                </observation>
              </entryRelationship>
            </observation>
          </entry>
          <entry>
            <procedure classCode="PROC" moodCode="RQO" >
            <!-- Consolidated Plan of Care Activity Procedure TemplateId (Implied Template) -->
              <templateId root="2.16.840.1.113883.10.20.22.4.41"/>
              <!-- QRDA Procedure, Order TemplateId -->
              <templateId root="2.16.840.1.113883.10.20.24.3.63"/>
              <id root="5106e039944dfe4d2000000d"/>
              <code code="313020008" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.1257.1.1533"><originalText>Procedure, Order: BH Electroconvulsive therapy (Code List: 2.16.840.1.113883.3.1257.1.1533)</originalText><translation code="90870" codeSystem="2.16.840.1.113883.6.12"/>
                <translation code="94.27" codeSystem="2.16.840.1.113883.6.103"/>
                <translation code="GZB4ZZZ" codeSystem="2.16.840.1.113883.6.4"/>
              </code>
              <text>Procedure, Order: BH Electroconvulsive therapy (Code List: 2.16.840.1.113883.3.1257.1.1533)</text>
              <statusCode code="new"/>
              <!-- Attribute: method-->
              <methodCode nullFlavor="UNK"/>
              <!-- Attribute: datetime -->
              <author>
                <time value='20110524094323'/>
                <assignedAuthor>
                 <id nullFlavor="NA"/>
                </assignedAuthor>
              </author>
            </procedure>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="GOL">
              <!-- 2.16.840.1.113883.10.20.22.4.44 Plan of Care Activity Observation -->
              <templateId root="2.16.840.1.113883.10.20.22.4.44"/>
              <templateId root="2.16.840.1.113883.10.20.24.3.1"/>
              <id root="F3D6FD73-B2C0-4274-BFD2-A485957734DB"/>
              <code code="252465000" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Pulse oximetry" sdtc:valueSet="2.16.840.1.113883.6.103"/>
              <text>Care Goal: Pulse Oximetry greater than 92%</text>
              <statusCode code="active"/>
              <effectiveTime>
                <low value="20110101140000"/>
              </effectiveTime>
            </observation>
          </entry>
          <entry>
            <!-- Patient characteristic expired -->
            <observation classCode="OBS" moodCode="EVN">
              <!-- Consolidation Deceased Observation templateId -->
              <templateId root="2.16.840.1.113883.10.20.22.4.79"/>
              <!-- Patient Characteristic Expired -->
              <templateId root="2.16.840.1.113883.10.20.24.3.54"/>
              <id root="22aeb750-4308-0130-0ade-680688cbd736"/>
              <code code="ASSERTION" codeSystem="2.16.840.1.113883.5.4" codeSystemName="HL7ActCode"/>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value="20151215010130" />
              </effectiveTime>
              <value code="419099009" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Dead" sdtc:valueSet="2.16.840.1.113883.3.666.5.539" xsi:type="CD"/>
            </observation>
          </entry>
          <entry>
            <!-- Clinical Trial Participant -->
            <observation classCode="OBS" moodCode="EVN">
              <templateId root="2.16.840.1.113883.10.20.24.3.51"/>
              <id root="22ab92c0-4308-0130-0ade-680688cbd736"/>
              <code code="ASSERTION" codeSystem="2.16.840.1.113883.5.4"/>
              <statusCode code="active"/>
              <effectiveTime>
                <low value="20100101000000"/> <!-- Currently hardcoded to January 1, 2010 -->
              </effectiveTime>
              <value code="428024001" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="clinical trial participant" xsi:type="CD"/>
            </observation>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="EVN" >
              <!-- Medication, Administered template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.42"/>
              <id root="50f84c187042f98775000108"/>
              <code code="416118004" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Administration"/>
              <statusCode code="completed"/>
              <effectiveTime xsi:type="IVL_TS">
                <low value='20061203201716'/>
                <high value='20061204072502'/>
              </effectiveTime>
              <entryRelationship typeCode="COMP">
                <substanceAdministration classCode="SBADM" moodCode="EVN">
                  <!-- Medication Activity (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.16"/>
                  <id root="278dade0-4307-0130-0add-680688cbd736"/>
                  <text>Medication, Administered: Pneumococcal Vaccine (Code List: 2.16.840.1.113883.3.464.1003.110.12.1027)</text>
                  <statusCode code="completed"/>
                  <effectiveTime xsi:type="IVL_TS">
                    <low value='20061203201716'/>
                    <high value='20061204072502'/>
                  </effectiveTime>
                  <consumable>
                    <manufacturedProduct classCode="MANU">
                      <!-- Medication Information (consolidation) template -->
                      <templateId root="2.16.840.1.113883.10.20.22.4.23"/>
                      <id root="278db0f0-4307-0130-0add-680688cbd736"/>
                      <manufacturedMaterial>
                        <code code="33" codeSystem="2.16.840.1.113883.6.59" sdtc:valueSet="2.16.840.1.113883.3.464.1003.110.12.1027"><originalText>Medication, Administered: Pneumococcal Vaccine (Code List: 2.16.840.1.113883.3.464.1003.110.12.1027)</originalText></code>
                      </manufacturedMaterial>
                    </manufacturedProduct>
                  </consumable>
                </substanceAdministration>
              </entryRelationship>
            </act>
          </entry>
          <entry>
            <!--Medication Order -->
            <substanceAdministration classCode="SBADM" moodCode="RQO" negationInd="true">
              <templateId root="2.16.840.1.113883.10.20.22.4.42"/>
              <!-- Medication, Order template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.47"/>
              <id root="50f84c1a7042f987750001d2"/>
              <text>Medication, Order: Beta Blocker Therapy for LVSD (Code List: 2.16.840.1.113883.3.526.3.1184)</text>
              <statusCode code="new"/>
              <effectiveTime xsi:type="IVL_TS">
                <low value='20000328001401'/>
                <high value='20000328012924'/>
              </effectiveTime>
              <consumable>
                <manufacturedProduct classCode="MANU">
                  <!-- Medication Information (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.23"/>
                  <id root="28c33d70-4307-0130-0add-680688cbd736"/>
                  <manufacturedMaterial>
                    <code code="866439" codeSystem="2.16.840.1.113883.6.88" sdtc:valueSet="2.16.840.1.113883.3.526.3.1184"><originalText>Medication, Order: Beta Blocker Therapy for LVSD (Code List: 2.16.840.1.113883.3.526.3.1184)</originalText></code>
                  </manufacturedMaterial>
                </manufacturedProduct>
              </consumable>
              <entryRelationship typeCode="RSON">
                <observation classCode="OBS" moodCode="EVN">
                  <templateId root="2.16.840.1.113883.10.20.24.3.88"/>
                  <code code="410666004"
                        codeSystem="2.16.840.1.113883.6.96"
                        displayName="reason"
                        codeSystemName="SNOMED CT"/>
                  <statusCode code="completed"/>
                  <effectiveTime value='20000328001401'/>
                  <value xsi:type="CD"
                       code="79899007"
                       codeSystem=""/>
                </observation>
              </entryRelationship>
            </substanceAdministration>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="EVN">
              <!-- Discharge Medication Entry -->
              <templateId root="2.16.840.1.113883.10.20.24.3.105"/>
              <id root="50f84dbb7042f9366f000189"/>
              <code code="10183-2" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" displayName="Discharge medication"/>
              <statusCode code="active"/>
              <effectiveTime>
                <low value="20050430111813"/>
              </effectiveTime>
              <entryRelationship typeCode="SUBJ">
                <substanceAdministration moodCode="EVN" classCode="SBADM">
                  <!-- Medication Activity (consolidation) template -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.16"/>
                  <!-- Medication, Active template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.41"/>
                  <id root="21305e00-4308-0130-0ade-680688cbd736"/>
                  <text>Medication, Discharge: Antithrombotic Therapy (Code List: 2.16.840.1.113883.3.117.1.7.1.201)</text>
                  <statusCode code="active"/>
                  <effectiveTime xsi:type="IVL_TS">
                    <low value="20050430111813"/>
                    <high value="20050501022146"/>
                  </effectiveTime>
                  <!-- Attribute: dose -->
                  <consumable>
                    <manufacturedProduct classCode="MANU">
                      <!-- Medication Information (consolidation) template -->
                      <templateId root="2.16.840.1.113883.10.20.22.4.23"/>
                      <id root="21306190-4308-0130-0ade-680688cbd736"/>
                      <manufacturedMaterial>
                        <code nullFlavor="UNK">
                          <originalText>Medication, Discharge: Antithrombotic Therapy (Code List: 2.16.840.1.113883.3.117.1.7.1.201)</originalText>
                          <translation code="994435" codeSystem="2.16.840.1.113883.6.88"/>
                        </code>
                      </manufacturedMaterial>
                      <manufacturerOrganization>
                        <name>Medication Factory Inc.</name>
                      </manufacturerOrganization>
                    </manufacturedProduct>
                  </consumable>
                </substanceAdministration>
              </entryRelationship>
            </act>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!-- consolidation CDA Allergy Observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.7"/>
              <templateId root="2.16.840.1.113883.10.20.24.3.46"/>
              <id root="50f84c1a7042f987750001db"/>
              <code code="ASSERTION"
                    displayName="Assertion"
                    codeSystem="2.16.840.1.113883.5.4"
                    codeSystemName="ActCode"/>
              <statusCode code="completed"/>
              <effectiveTime>
                  <low value="20061203201716"/>
              </effectiveTime>
              <value xsi:type="CD"
                     code="59037007"
                     displayName="Drug intolerance"
                     codeSystem="2.16.840.1.113883.6.96"
                     codeSystemName="SNOMED CT"/>

              <participant typeCode="CSM">
                <participantRole classCode="MANU">
                  <playingEntity classCode="MMAT">
                    <code code="998695" codeSystem="2.16.840.1.113883.6.88" sdtc:valueSet="2.16.840.1.113883.3.526.3.1174"><originalText>Medication, Intolerance: Beta Blocker Therapy (Code List: 2.16.840.1.113883.3.526.3.1174)</originalText></code>
                    <text>Medication, Intolerance: Beta Blocker Therapy (Code List: 2.16.840.1.113883.3.526.3.1174)</text>
                  </playingEntity>
                </participantRole>
              </participant>
            </observation>
          </entry>
          <entry>
            <observation classCode="OBS" moodCode="EVN">
              <!-- consolidation CDA Allergy observation template -->
              <templateId root="2.16.840.1.113883.10.20.22.4.7"/>
              <!--  Medication Allergy -->
              <templateId root="2.16.840.1.113883.10.20.24.3.44"/>
              <id root="50f84db97042f9366f00000e"/>
              <code code="ASSERTION" displayName="Assertion" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='19790809140056'/>
              </effectiveTime>
              <value code="62014003" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Adverse drug effect" xsi:type="CD"/>

              <participant typeCode="CSM">
                <participantRole classCode="MANU">
                  <playingEntity classCode="MMAT">
                    <code code="996994" codeSystem="2.16.840.1.113883.6.88" sdtc:valueSet="2.16.840.1.113883.3.666.5.626"><originalText>Medication, Adverse Effects: Hospital Measures-Aspirin (Code List: 2.16.840.1.113883.3.666.5.626)</originalText></code>
                    <name>Medication, Adverse Effects: Hospital Measures-Aspirin (Code List: 2.16.840.1.113883.3.666.5.626)</name>
                  </playingEntity>
                </participantRole>
              </participant>
            </observation>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="EVN" >
              <!-- Communication from provider to provider -->
              <templateId root="2.16.840.1.113883.10.20.24.3.4"/>
              <id root="50f84c1d7042f987750003bf"/>
              <code code="371545006" codeSystem="2.16.840.1.113883.6.96" sdtc:valueSet="2.16.840.1.113883.3.464.1003.121.12.1006"><originalText>Communication: From Provider to Provider: Consultant Report (Code List: 2.16.840.1.113883.3.464.1003.121.12.1006)</originalText></code>
              <text>Communication: From Provider to Provider: Consultant Report (Code List: 2.16.840.1.113883.3.464.1003.121.12.1006)</text>
              <statusCode code="completed"/>
              <effectiveTime>
                <low value='19810627142601'/>
              </effectiveTime>
              <participant typeCode="AUT">
                <participantRole classCode="ASSIGNED">
                  <code code="158965000" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Medical Practitioner"/>
                </participantRole>
              </participant>

              <participant typeCode="IRCP">
                <participantRole classCode="ASSIGNED">
                  <code code="158965000" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Medical Practitioner"/>
                </participantRole>
              </participant>
            </act>
          </entry>
          <entry>
            <act classCode="ACT" moodCode="EVN">
              <templateId root="2.16.840.1.113883.10.20.24.3.3"/>
              <id root="50cf48409eae47465700008f"/>
              <code nullFlavor="UNK" ><originalText>Communication: From Provider to Patient: Asthma Management Plan (Code List: 2.16.840.1.113883.3.117.1.7.1.131)</originalText><translation code="69981-9" codeSystem="2.16.840.1.113883.6.1"/>
              </code>
              <text>Communication: From Provider to Patient: Asthma Management Plan (Code List: 2.16.840.1.113883.3.117.1.7.1.131)</text>
              <statusCode code="completed"/>
              <effectiveTime>
              <low value='20100605220000'/>
              </effectiveTime>
              <participant typeCode="AUT">
                <participantRole classCode="ASSIGNED">
                  <code code="158965000" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT" displayName="Medical Practitioner"/>
                </participantRole>
              </participant>
              <participant typeCode="IRCP">
                <participantRole classCode="PAT"/>
              </participant>
            </act>
          </entry>

          <entry>
            <act classCode="ACT" moodCode="EVN">
              <!-- Conforms to C-CDA 2.1 Problem Concern Act (V3) -->
              <templateId root="2.16.840.1.113883.10.20.22.4.3" extension="2015-08-01" />
              <!-- Diagnosis Concern Act -->
              <templateId root="2.16.840.1.113883.10.20.24.3.137" />
              <id root="50cf48409eae47465700008f" />
              <code code="CONC" codeSystem="2.16.840.1.113883.5.6" displayName="Concern" />
              <statusCode code="active" />
              <effectiveTime>
                <!-- Attribute: admission datetime -->
                <low value="19890903081502"/>
                <!-- Attribute: discharge datetime -->
                <high value="19890904034509"/>
              </effectiveTime>
              <entryRelationship typeCode="SUBJ">
                <observation classCode="OBS" moodCode="EVN">
                  <!-- Conforms to C-CDA R2.1 Problem Observation (V3) -->
                  <templateId root="2.16.840.1.113883.10.20.22.4.4" extension="2015-08-01" />
                  <!-- Diagnosis template -->
                  <templateId root="2.16.840.1.113883.10.20.24.3.135" />
                  <id root="50f84db97042f9366f00001b"/>
                  <code code="282291009" displayName="diagnosis" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED-CT"/>
                  <text>Diagnosis Hospital Measures - Any infection (Code List: 2.16.840.1.113883.3.666.5.696)</text>
                  <statusCode code="completed"/>
                  <effectiveTime>
                    <!-- Attribute: admission datetime -->
                    <low value="19890903081502"/>
                    <!-- Attribute: discharge datetime -->
                    <high value="19890904034509"/>
                  </effectiveTime>
                  <value code="999.34" codeSystem="2.16.840.1.113883.6.103" xsi:type="CD" sdtc:valueSet="2.16.840.1.113883.3.666.5.696">
                    <originalText>Diagnosis Hospital Measures - Any infection (Code List: 2.16.840.1.113883.3.666.5.696)</originalText>
                    <translation code="T86.842" codeSystem="2.16.840.1.113883.6.90"/>
                  </value>
                  <entryRelationship typeCode="REFR">
                    <observation classCode="OBS" moodCode="EVN">
                      <!-- Problem Status (consolidation) template -->
                      <templateId root="2.16.840.1.113883.10.20.22.4.6"/>
                      <!-- Problem Status, Active template -->
                      <templateId root="2.16.840.1.113883.10.20.24.3.94"/>
                      <id root="1fddfe00-4308-0130-0ade-680688cbd736"/>
                      <code code="33999-4" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" displayName="status"/>
                      <statusCode code="completed"/>
                      <value xsi:type="CD" code="55561003" displayName="active" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT"/>
                    </observation>
                  </entryRelationship>

                 <entryRelationship typeCode="REFR">
                    <observation classCode="OBS" moodCode="EVN">
                      <!-- Problem Status (consolidation) template -->
                      <templateId root="2.16.840.1.113883.10.20.22.4.6"/>
                      <!-- Problem Status, Active template -->
                      <templateId root="2.16.840.1.113883.10.20.24.3.94"/>
                      <id root="1fddfe00-4308-0130-0ade-680688cbd736"/>
                      <code code="33999-4" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" displayName="status"/>
                      <statusCode code="completed"/>
                      <value xsi:type="CD" code="55561003" displayName="active" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT"/>
                    </observation>
                  </entryRelationship>

                  <entryRelationship typeCode="REFR">
                   <observation classCode="OBS" moodCode="EVN">
                      <templateId root="2.16.840.1.113883.10.20.22.4.8"/>
                      <code code="SEV"
                            codeSystem="2.16.840.1.113883.5.4"/>
                       <statusCode code="completed"/>
                       <value xsi:type="CD" code="55561003" displayName="active" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED CT"/>
                    </observation>

                   </entryRelationship>
                 </observation>
              </entryRelationship>
            </act>
          </entry>

          <entry>
            <act classCode="ACT" moodCode="EVN" >
              <!-- Immunization, Administered template -->
              <templateId root="2.16.840.1.113883.10.20.24.3.140"/>
              <id root="9c09831a-9e99-4f9c-96de-9ce334f943d7"/>
              <code code="416118004" displayName="Administration" codeSystem="2.16.840.1.113883.6.96" codeSystemName="SNOMED-CT"/>
              <statusCode code="completed"/>
              <effectiveTime xsi:type="IVL_TS">
                  <low value='19890509170647'/>
                  <high value='19890509173724'/>
              </effectiveTime>
              <entryRelationship typeCode="COMP">
                <substanceAdministration classCode="SBADM" moodCode="EVN">
                    <!-- C-CDA R2 Immunization Activity -->
                    <templateId root="2.16.840.1.113883.10.20.22.4.52" extension="2014-06-09"/>
                    <id root="1ea77951-61f4-425f-bb69-dd396d63c44f"/>
                    <statusCode code="completed"/>
                    <effectiveTime xsi:type="IVL_TS">
                      <low value='19890509170647'/>
                      <high value='19890509173724'/>
                    </effectiveTime>
                    <consumable>
                      <manufacturedProduct classCode="MANU">
                        <!-- C-CDA R2.1 Immunization Medication Information (V2) -->
                        <templateId root="2.16.840.1.113883.10.20.22.4.23"/>
                        <id root="28c6a620-4307-0130-0add-680688cbd736"/>
                        <manufacturedMaterial>
                          <code code="33" codeSystem="2.16.840.1.113883.6.59" displayName="Pneumococcal polysaccharide vaccine" codeSystemName="CVX"/>
                        </manufacturedMaterial>
                        <manufacturerOrganization>
                          <name>Medication Factory Inc.</name>
                        </manufacturerOrganization>
                      </manufacturedProduct>
                    </consumable>
                </substanceAdministration>
              </entryRelationship>
            </act>
          </entry>
        </section>
      </component>
    </structuredBody>
  </component>
</ClinicalDocument>
`)
